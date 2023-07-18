package sdk

import (
	"errors"
	"fmt"
	"github.com/alanshaw/go-carbites"
	"github.com/kataras/golog"
	"github.com/paradeum-team/chainstorage-sdk/code"
	"github.com/paradeum-team/chainstorage-sdk/consts"
	"github.com/paradeum-team/chainstorage-sdk/model"
	"github.com/paradeum-team/chainstorage-sdk/utils"
	"github.com/ulule/deepcopier"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type Upload struct {
	Config *Configuration
	Client *RestyClient
	logger *golog.Logger
}

// sema is a counting semaphore for limiting concurrency in dirEntries.
var sema = make(chan struct{}, 20)

func (u *Upload) UploadData(bucketName, dataPath string) (model.ObjectCreateResponse, error) {
	response := model.ObjectCreateResponse{}

	// 桶名称
	if err := checkBucketName(bucketName); err != nil {
		return response, err
	}

	// 上传对象路径
	if _, err := os.Stat(dataPath); os.IsNotExist(err) {
		return response, code.ErrCarUploadFileInvalidDataPath
	} else if err != nil {
		return response, err
	}

	// 确认桶数据有效性
	bucket := Bucket{Config: u.Config, Client: u.Client, logger: u.logger}
	respBucket, err := bucket.GetBucketByName(bucketName)
	if err != nil {
		return response, err
	}

	statusCode := int(respBucket.Code)
	if statusCode != http.StatusOK {
		return response, errors.New(respBucket.Msg)
	}

	// 桶ID
	bucketId := respBucket.Data.Id

	// 检查上传数据使用限制
	storageNetworkCode := respBucket.Data.StorageNetworkCode
	err = u.checkDataUsageLimitation(storageNetworkCode, dataPath)
	if err != nil {
		return response, err
	}

	// 对象上传
	response, err = u.uploadCarFile(bucketId, dataPath)
	if err != nil {
		return response, err
	}

	return response, nil
}

// 上传CAR文件
func (u *Upload) uploadCarFile(bucketId int, dataPath string) (model.ObjectCreateResponse, error) {
	response := model.ObjectCreateResponse{}

	// 数据路径为空
	if len(dataPath) == 0 {
		return response, code.ErrCarUploadFileInvalidDataPath
	}

	// 数据路径无效
	fileInfo, err := os.Stat(dataPath)
	if os.IsNotExist(err) {
		return response, code.ErrCarUploadFileInvalidDataPath
	} else if err != nil {
		//log.WithError(err).WithField("dataPath", dataPath).Error("fail to return stat of file")
		u.logger.Errorf(fmt.Sprintf("fail to return stat of file, method:uploadCarFile, dataPath:%s, err:%+v\n", dataPath, err))
		return response, err
	}

	carFileUploadReq := model.CarFileUploadReq{}
	// 上传为目录的情况
	if fileInfo.IsDir() {
		notEmpty, err := u.isFolderNotEmpty(dataPath)
		if err != nil {
			//log.WithError(err).WithField("dataPath", dataPath).Error("fail to check uploading folder")
			u.logger.Errorf(fmt.Sprintf("fail to check uploading folder, method:uploadCarFile, dataPath:%s, err:%+v\n", dataPath, err))
			return response, err
		}

		if !notEmpty {
			return response, code.ErrCarUploadFileInvalidDataFolder
		}

		carFileUploadReq.ObjectTypeCode = consts.ObjectTypeCodeDir
	}

	car := Car{Config: u.Config, Client: u.Client, logger: u.logger}
	fileDestination := car.GenerateTempFileName(utils.CurrentDate()+"_", ".tmp")
	carVersion := u.Config.CarVersion

	//log.WithFields(logrus.Fields{
	//	"fileDestination": fileDestination,
	//	"dataPath":        dataPath,
	//	"carVersion":      carVersion,
	//	"begintime":       GetTimestampString(),
	//}).Info("Create car file start")
	u.logger.Infof("Create car file start, begintime:%s, carVersion:%d, fileDestination:%s, dataPath:%s\n", u.getTimestampString(), carVersion, fileDestination, dataPath)
	// 创建Car文件
	err = car.CreateCarFile(dataPath, fileDestination)
	if err != nil {
		//log.WithError(err).
		//	WithFields(logrus.Fields{
		//		"fileDestination": fileDestination,
		//		"dataPath":        dataPath,
		//	}).Error("Fail to create car file")
		u.logger.Errorf("Fail to create car file, carVersion:%d, fileDestination:%s, dataPath:%s, error:%+v\n", carVersion, fileDestination, dataPath, err)
		return response, code.ErrCarUploadFileCreateCarFileFail
	}
	//log.WithFields(logrus.Fields{
	//	"fileDestination": fileDestination,
	//	"dataPath":        dataPath,
	//	"carVersion":      carVersion,
	//	"endtime":         GetTimestampString(),
	//}).Info("Create car file finish")
	u.logger.Infof("Create car file finish, endtime:%s, carVersion:%d, fileDestination:%s, dataPath:%s\n", u.getTimestampString(), carVersion, fileDestination, dataPath)

	defer func(fileDestination string) {
		// todo: add confg to control deleting temp files?
		//if !u.Config.CleanTmpData {
		//	return
		//}

		err := os.Remove(fileDestination)
		if err != nil {
			//log.WithError(err).
			//	WithFields(logrus.Fields{
			//		"fileDestination": fileDestination,
			//	}).Error("Fail to remove car file")
			u.logger.Errorf("Fail to remove car file, fileDestination:%s, error:%+v\n", fileDestination, err)
		}
	}(fileDestination)

	// 解析CAR文件，获取DAG信息，获取文件或目录的CID
	rootLink := model.RootLink{}
	err = car.ParseCarFile(fileDestination, &rootLink)
	if err != nil {
		//log.WithError(err).
		//	WithFields(logrus.Fields{
		//		"fileDestination": fileDestination,
		//	}).Error("Fail to parse car file")
		u.logger.Errorf("Fail to parse car file, fileDestination:%s, error:%+v\n", fileDestination, err)
		return response, code.ErrCarUploadFileParseCarFileFail
	}

	rootCid := rootLink.RootCid.String()
	objectCid := rootLink.Cid.String()
	objectSize := int64(rootLink.Size)
	objectName := rootLink.Name

	// 设置请求参数
	carFileUploadReq.BucketId = bucketId
	carFileUploadReq.ObjectCid = objectCid
	carFileUploadReq.ObjectSize = objectSize
	carFileUploadReq.ObjectName = objectName
	carFileUploadReq.FileDestination = fileDestination
	carFileUploadReq.CarFileCid = rootCid

	// 计算文件sha256
	sha256, err := utils.GetFileSha256ByPath(fileDestination)
	if err != nil {
		//log.WithError(err).
		//	WithFields(logrus.Fields{
		//		"fileDestination": fileDestination,
		//	}).Error("Fail to calculate file sha256")
		u.logger.Errorf("Fail to calculate file sha256, fileDestination:%s, error:%+v\n", fileDestination, err)
		return response, code.ErrCarUploadFileComputeCarFileHashFail
	}
	carFileUploadReq.RawSha256 = sha256

	// 使用Root CID秒传检查
	object := Object{Config: u.Config, Client: u.Client, logger: u.logger}
	objectExistResponse, err := object.IsExistObjectByCid(objectCid)
	if err != nil {
		//log.WithError(err).
		//	WithFields(logrus.Fields{
		//		"objectCid":           objectCid,
		//		"objectExistResponse": objectExistResponse,
		//	}).Error("Fail to check if exist object")
		u.logger.Errorf("Fail to check if exist object, objectCid:%s, objectExistResponse:%+v, error:%+v\n", objectCid, objectExistResponse, err)
		return response, code.ErrCarUploadFileReferenceObjcetFail
	}

	// CID存在，执行秒传操作
	objectExistCheck := objectExistResponse.Data
	if objectExistCheck.IsExist {
		response, err := car.ReferenceObject(&carFileUploadReq)
		if err != nil {
			//log.WithError(err).
			//	WithFields(logrus.Fields{
			//		"carFileUploadReq": carFileUploadReq,
			//		"response":         response,
			//	}).Error("Fail to reference object")
			u.logger.Errorf("Fail to reference object, carFileUploadReq:%+v, response:%+v, error:%+v\n", carFileUploadReq, response, err)
			return response, code.ErrCarUploadFileReferenceObjcetFail
		}

		return response, nil
	}

	// CAR文件大小，超过分片阈值
	carFileInfo, err := os.Stat(fileDestination)
	if err != nil {
		//log.WithError(err).WithField("fileDestination", fileDestination).Error("fail to return stat of file")
		u.logger.Errorf("fail to return stat of file, fileDestination:%s, error:%+v\n", fileDestination, err)
		return response, err
	}

	carFileSize := carFileInfo.Size()
	carFileShardingThreshold := u.Config.CarFileShardingThreshold

	// 生成CAR分片文件上传
	if carFileSize > int64(carFileShardingThreshold) {
		response, err = u.uploadBigCarFile(&carFileUploadReq)
		if err != nil {
			//log.WithError(err).
			//	WithFields(logrus.Fields{
			//		"carFileUploadReq": carFileUploadReq,
			//		"response":         response,
			//	}).Error("Fail to upload big car file")
			u.logger.Errorf("Fail to upload big car file, carFileUploadReq:%+v, response:%+v, error:%+v\n", carFileUploadReq, response, err)
			return response, code.ErrCarUploadFileFail
		}

		return response, nil
	}

	// 普通上传
	file, err := os.Open(fileDestination)
	if err != nil {
		//log.WithError(err).WithField("fileDestination", fileDestination).Error("fail to return stat of file")
		u.logger.Errorf("fail to return stat of file, fileDestination:%s, error:%+v\n", fileDestination, err)
		return response, err
	}
	defer file.Close()

	//log.WithFields(logrus.Fields{
	//	"carFileUploadReq": carFileUploadReq,
	//	"begintime":        GetTimestampString(),
	//}).Info("Upload car file start")
	u.logger.Infof("Upload car file start, begintime:%s, carFileUploadReq:%+v\n", u.getTimestampString(), carFileUploadReq)
	response, err = car.UploadCarFile(&carFileUploadReq)
	if err != nil {
		//log.WithError(err).
		//	WithFields(logrus.Fields{
		//		"carFileUploadReq": carFileUploadReq,
		//		"response":         response,
		//	}).Error("Fail to upload car file")
		u.logger.Errorf("Fail to upload car file, carFileUploadReq:%+v, response:%+v, error:%+v\n", carFileUploadReq, response, err)
		return response, code.ErrCarUploadFileFail
	}
	//log.WithFields(logrus.Fields{
	//	"response": response,
	//	"endtime":  GetTimestampString(),
	//}).Info("Upload car file finish")
	u.logger.Infof("Upload car file finish, endtime:%s, response:%+v\n", u.getTimestampString(), response)

	return response, err
}

// 上传大CAR文件
func (u *Upload) uploadBigCarFile(req *model.CarFileUploadReq) (model.ObjectCreateResponse, error) {
	response := model.ObjectCreateResponse{}

	//log.WithFields(logrus.Fields{
	//	"req":       req,
	//	"begintime": GetTimestampString(),
	//}).Info("Generate sharding car files start")
	u.logger.Infof("Generate sharding car files start, begintime:%s, req:%+v\n", u.getTimestampString(), req)
	// 生成CAR分片文件
	//shardingCarFileUploadReqs := []model.CarFileUploadReq{}
	var shardingCarFileUploadReqs []model.CarFileUploadReq
	//err := car.GenerateShardingCarFiles(req, &shardingCarFileUploadReqs)
	err := u.generateShardingCarFiles(req, &shardingCarFileUploadReqs)
	if err != nil {
		//log.WithError(err).
		//	WithFields(logrus.Fields{
		//		"req":                       req,
		//		"shardingCarFileUploadReqs": shardingCarFileUploadReqs,
		//	}).Error("Fail to generate sharding car files")
		u.logger.Errorf("Fail to generate sharding car files, req:%+v, shardingCarFileUploadReqs:%+v, error:%+v\n", req, shardingCarFileUploadReqs, err)
		return response, err
	}
	//log.WithFields(logrus.Fields{
	//	"shardingCarFileUploadReqs": shardingCarFileUploadReqs,
	//	"endtime":                   GetTimestampString(),
	//}).Info("Generate sharding car files finish")
	u.logger.Infof("Generate sharding car files finish, endtime:%s, shardingCarFileUploadReqs:%+v\n", u.getTimestampString(), shardingCarFileUploadReqs)

	// 删除CAR分片文件
	defer func(shardingCarFileUploadReqs []model.CarFileUploadReq) {
		// todo: add config to control deleting temp files?
		//if !cliConfig.CleanTmpData {
		//	return
		//}

		for i := range shardingCarFileUploadReqs {
			fileDestination := shardingCarFileUploadReqs[i].FileDestination
			err := os.Remove(fileDestination)
			if err != nil {
				//log.WithError(err).
				//	WithFields(logrus.Fields{
				//		"fileDestination": fileDestination,
				//	}).Error("Fail to remove sharding car file")
				u.logger.Errorf("Fail to remove sharding car file, fileDestination:%s, error:%+v\n", fileDestination, err)
			}
		}
	}(shardingCarFileUploadReqs)

	// 计算总文件大小
	totalSize := int64(0)
	for i := range shardingCarFileUploadReqs {
		totalSize += shardingCarFileUploadReqs[i].ObjectSize
	}

	// 上传CAR文件分片
	car := Car{Config: u.Config, Client: u.Client, logger: u.logger}
	maxRetries := 3
	retryDelay := time.Duration(3) * time.Second

	//log.WithFields(logrus.Fields{
	//	"shardingCarFileUploadReqs": shardingCarFileUploadReqs,
	//	"begintime":                 GetTimestampString(),
	//}).Info("Upload sharding car files start")
	u.logger.Infof("Upload sharding car files start, begintime:%s, shardingCarFileUploadReqs:%+v\n", u.getTimestampString(), shardingCarFileUploadReqs)
	//uploadRespList := []model.ShardingCarFileUploadResponse{}
	var uploadRespList []model.ShardingCarFileUploadResponse
	for i := range shardingCarFileUploadReqs {
		for j := 0; j < maxRetries; j++ {
			uploadingReq := model.CarFileUploadReq{}
			deepcopier.Copy(&shardingCarFileUploadReqs[i]).To(&uploadingReq)

			//log.WithFields(logrus.Fields{
			//	"uploadingReq": uploadingReq,
			//	"index":        i,
			//	"retry":        j,
			//}).Info("upload sharding car file")
			u.logger.Infof("upload sharding car file, uploadingReq:%+v, index:%d, retry:%d\n", uploadingReq, i, j)
			uploadResp, err := car.UploadShardingCarFile(&uploadingReq)
			if err == nil && uploadResp.Code == http.StatusOK {
				uploadRespList = append(uploadRespList, uploadResp)
				break
			}

			// 记录日志
			//fmt.Printf("UploadBigCarFile => UploadShardingCarFileExt, index:%d, uploadResp:%+v\n", i, uploadResp)
			//log.WithError(err).
			//	WithFields(logrus.Fields{
			//		"uploadingReq": uploadingReq,
			//		"uploadResp":   uploadResp,
			//		"index":        i,
			//		"retry":        j,
			//	}).Error("Fail to upload sharding car file")
			u.logger.Errorf("Fail to upload sharding car file, uploadingReq:%+v, uploadResp:%+v, index:%d, retry:%d, error:%+v\n", uploadingReq, uploadResp, i, j, err)

			if j == maxRetries-1 {
				// 尝试maxRetries次失败
				if err != nil {
					return response, err
				} else if uploadResp.Code != http.StatusOK {
					return response, errors.New(response.Msg)
				}
			}

			time.Sleep(retryDelay)
		}
	}
	//log.WithFields(logrus.Fields{
	//	"shardingCarFileUploadReqs": shardingCarFileUploadReqs,
	//	"endtime":                   GetTimestampString(),
	//}).Info("Upload sharding car files finish")
	u.logger.Infof("Upload sharding car files finish, endtime:%s, shardingCarFileUploadReqs:%+v\n", u.getTimestampString(), shardingCarFileUploadReqs)

	//log.WithFields(logrus.Fields{
	//	"req":       req,
	//	"begintime": GetTimestampString(),
	//}).Info("Confirm sharding car files start")
	u.logger.Infof("Confirm sharding car files start, begintime:%s, req:%+v\n", u.getTimestampString(), req)
	// 确认分片上传成功
	response, err = car.ConfirmShardingCarFiles(req)
	if err != nil {
		//log.WithError(err).
		//	WithFields(logrus.Fields{
		//		"req":      req,
		//		"response": response,
		//	}).Error("Fail to Confirm sharding car files")
		u.logger.Errorf("Fail to Confirm sharding car files, req:%+v, response:%+v, error:%+v\n", req, response, err)
		return response, err
	}
	//log.WithFields(logrus.Fields{
	//	"response": response,
	//	"endtime":  GetTimestampString(),
	//}).Info("Confirm sharding car files finish")
	u.logger.Infof("Confirm sharding car files finish, endtime:%s, response:%+v\n", u.getTimestampString(), response)

	return response, nil
}

// 生成CAR分片文件
func (u *Upload) generateShardingCarFiles(req *model.CarFileUploadReq, shardingCarFileUploadReqs *[]model.CarFileUploadReq) error {
	fileDestination := req.FileDestination

	bigCarFile, err := os.Open(fileDestination)
	if err != nil {
		//log.WithError(err).
		//	WithFields(logrus.Fields{
		//		"fileDestination": fileDestination,
		//	}).Error("Fail to open car file")
		u.logger.Errorf("Fail to open car file, fileDestination:%s, error:%+v\n", fileDestination, err)
		return err
	}
	defer bigCarFile.Close()

	// CAR文件分片设置
	targetSize := u.Config.CarFileShardingThreshold
	splitter, _ := carbites.NewTreewalkSplitterFromPath(fileDestination, targetSize)

	shardingNo := 1
	for {
		car, err := splitter.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			//log.WithError(err).
			//	WithFields(logrus.Fields{
			//		"shardingNo": shardingNo,
			//	}).Error("Fail to generate sharding car file")
			u.logger.Errorf("Fail to generate sharding car file, shardingNo:%d, error:%+v\n", shardingNo, err)
			return err
		}

		bytes, err := io.ReadAll(car)
		if err != nil {
			//log.WithError(err).
			//	WithFields(logrus.Fields{
			//		"shardingNo": shardingNo,
			//	}).Error("Fail to generate sharding car file")
			u.logger.Errorf("Fail to generate sharding car file, shardingNo:%d, error:%+v\n", shardingNo, err)
			return err
		}

		// 设置文件名称
		filename := fmt.Sprintf("_chunk.c%d", shardingNo)
		shardingFileDestination := strings.Replace(fileDestination, filepath.Ext(fileDestination), filename, 1)

		chunkSize := int64(len(bytes))

		// 生成分片文件
		err = os.WriteFile(shardingFileDestination, bytes, 0644)
		if err != nil {
			//log.WithError(err).
			//	WithFields(logrus.Fields{
			//		"shardingNo":              shardingNo,
			//		"shardingFileDestination": shardingFileDestination,
			//	}).Error("Fail to generate sharding car file")
			u.logger.Errorf("Fail to generate sharding car file, shardingNo:%d, shardingFileDestination:%s, error:%+v\n", shardingNo, shardingFileDestination, err)
			return err
		}

		// 计算分片文件sha256
		shardingSha256, err := utils.GetFileSha256ByPath(shardingFileDestination)
		if err != nil {
			//log.WithError(err).
			//	WithFields(logrus.Fields{
			//		"shardingNo":              shardingNo,
			//		"shardingFileDestination": shardingFileDestination,
			//	}).Error("Fail to calculate file sha256")
			u.logger.Errorf("Fail to calculate file sha256, shardingNo:%d, shardingFileDestination:%s, error:%+v\n", shardingNo, shardingFileDestination, err)
			return err
		}

		// 设置分片请求对象
		shardingCarFileUploadReq := model.CarFileUploadReq{}
		deepcopier.Copy(req).To(&shardingCarFileUploadReq)
		shardingCarFileUploadReq.FileDestination = shardingFileDestination
		shardingCarFileUploadReq.ShardingSha256 = shardingSha256
		shardingCarFileUploadReq.ShardingNo = shardingNo
		shardingCarFileUploadReq.ObjectSize = chunkSize
		*shardingCarFileUploadReqs = append(*shardingCarFileUploadReqs, shardingCarFileUploadReq)

		shardingNo++
	}

	// 分片失败
	shardingAmount := len(*shardingCarFileUploadReqs)
	if shardingAmount == 0 {
		//log.WithError(err).
		//	WithFields(logrus.Fields{
		//		"shardingCarFileUploadReqs": shardingCarFileUploadReqs,
		//	}).Error("Fail to generate sharding car file")
		u.logger.Errorf("Fail to generate sharding car file, shardingCarFileUploadReqs:%+v, error:%+v\n", shardingCarFileUploadReqs, err)
		return code.ErrCarUploadFileChunkCarFileFail
	}

	req.ShardingAmount = shardingAmount

	return nil
}

// region auxiliary method

// 检查上传数据使用限制
func (u *Upload) checkDataUsageLimitation(storageNetworkCode int, dataPath string) error {
	bucket := Bucket{Config: u.Config, Client: u.Client, logger: u.logger}

	// 检查可用空间
	usersQuotaResp, err := bucket.GetUsersQuotaByStorageNetworkCode(storageNetworkCode)
	if err != nil {
		return err
	}

	usersQuota := usersQuotaResp.Data
	usersQuotaDetails := usersQuota.Details
	if len(usersQuotaDetails) == 0 {
		return code.ErrBucketQuotaFetchFail
	}

	// 基础版本
	isBasicVersion := usersQuota.PackagePlanId == 21001
	availableStorageSpace := int64(0)
	availableFileAmount := int64(0)
	availableUploadDirItems := int64(0)

	for _, usersQuotaDetail := range usersQuotaDetails {
		// 空间存储限制
		if usersQuotaDetail.ConstraintName == consts.ConstraintStorageSpace.String() {
			//availableStorageSpace = usersQuotaDetail.Available
			availableStorageSpace = usersQuotaDetail.LimitedQuota - usersQuotaDetail.UsedQuota
		}

		// 对象存储限制
		if usersQuotaDetail.ConstraintName == consts.ConstraintFileLimited.String() {
			//availableFileAmount = usersQuotaDetail.Available
			availableFileAmount = usersQuotaDetail.LimitedQuota - usersQuotaDetail.UsedQuota
		}

		// 上传文件夹条目限制
		if usersQuotaDetail.ConstraintName == consts.ConstraintUploadDirItems.String() {
			//availableUploadDirItems = usersQuotaDetail.Available
			availableUploadDirItems = usersQuotaDetail.LimitedQuota - usersQuotaDetail.UsedQuota
		}
	}

	// 可用文件存储限制超限
	if availableFileAmount <= 0 {
		return code.ErrCarUploadFileExccedObjectAmountUsage
	}

	// 获取上传数据使用量
	fileAmount, totalSize, err := u.getUploadingDataUsage(dataPath)
	if err != nil {
		return err
	}

	// 上传文件夹条目限制超限
	if fileAmount > availableUploadDirItems {
		return code.ErrCarUploadFileExccedObjectAmountUsage
	}

	if isBasicVersion {
		// 可用存储空间超限
		if totalSize > availableStorageSpace {
			return code.ErrCarUploadFileExccedStorageSpaceUsage
		}
	}

	return nil
}

// 获取上传数据使用量
func (u *Upload) getUploadingDataUsage(dataPath string) (int64, int64, error) {
	var totalSize int64
	var fileAmount int64

	// 数据路径为空
	if len(dataPath) == 0 {
		return 0, 0, code.ErrCarUploadFileInvalidDataPath
	}

	// 数据路径无效
	fileInfo, err := os.Stat(dataPath)
	if os.IsNotExist(err) {
		return 0, 0, code.ErrCarUploadFileInvalidDataPath
	} else if err != nil {
		//log.WithError(err).
		//	WithField("dataPath", dataPath).
		//	Error("fail to return stat of file")
		u.logger.Errorf(fmt.Sprintf("fail to return stat of file, method:getUploadingDataUsage, dataPath:%s, err:%+v\n", dataPath, err))
		return 0, 0, err
	}

	if !fileInfo.IsDir() {
		fileAmount++
		totalSize = fileInfo.Size()
		return fileAmount, totalSize, nil
	}

	fileSizes := make(chan int64)
	var wg sync.WaitGroup
	wg.Add(1)
	go u.walkDir(dataPath, &wg, fileSizes)

	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	for {
		size, ok := <-fileSizes
		if !ok {
			break // fileSizes was closed
		}

		fileAmount++
		totalSize += size
	}

	return fileAmount, totalSize, nil
}

func (u *Upload) walkDir(dir string, wg *sync.WaitGroup, fileSizes chan<- int64) {
	defer wg.Done()
	for _, entry := range u.dirEntries(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subDir := filepath.Join(dir, entry.Name())
			go u.walkDir(subDir, wg, fileSizes)
		} else {
			fileInfo, err := entry.Info()
			if err != nil {
				//log.WithError(err).
				//	WithField("dir", dir).
				//	Error("fail to return stat of file")
				u.logger.Errorf(fmt.Sprintf("fail to return stat of file, method:walkDir, dir:%s, err:%+v\n", dir, err))
				return
			}

			fileSizes <- fileInfo.Size()
		}
	}
}

// dirEntries returns the entries of directory dir.
func (u *Upload) dirEntries(dir string) []os.DirEntry {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token

	entries, err := os.ReadDir(dir)
	if err != nil {
		//log.WithError(err).
		//	WithField("dir", dir).
		//	Error("fail to read dir")
		u.logger.Errorf(fmt.Sprintf("fail to read dir, method:dirEntries, dir:%s, err:%+v\n", dir, err))
		return nil
	}

	return entries
}

func (u *Upload) isFolderNotEmpty(path string) (bool, error) {
	// Check if the path is a directory
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	if !fileInfo.IsDir() {
		return false, nil
	}

	// Open the directory
	dir, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer dir.Close()

	// Read the directory entries
	_, err = dir.Readdirnames(1)
	if err == nil {
		// Directory is not empty
		return true, nil
	} else if err == io.EOF {
		// Directory is empty
		return false, nil
	} else {
		// An error occurred while reading the directory
		return false, err
	}
}

func (u *Upload) getTimestampString() string {
	timestampString := time.Now().Format("2006-01-02 15:04:05.000000000") //当前时间的字符串，2006-01-02 15:04:05据说是golang的诞生时间，固定写法
	return timestampString
}

// endregion auxiliary method
