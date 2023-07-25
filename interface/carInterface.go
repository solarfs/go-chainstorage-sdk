package _interface

import (
	"github.com/solarfs/go-chainstorage-sdk/model"
	"io"
)

type Car interface {
	CreateCarFile(dataPath string, fileDestination string) error
	SplitCarFile(carFilePath string, chunkedFileDestinations *[]string) error
	ReferenceObject(req *model.CarFileUploadReq) (model.ObjectCreateResponse, error)
	UploadCarFile(req *model.CarFileUploadReq) (model.ObjectCreateResponse, error)
	UploadShardingCarFile(req *model.CarFileUploadReq) (model.ShardingCarFileUploadResponse, error)
	ConfirmShardingCarFiles(req *model.CarFileUploadReq) (model.ObjectCreateResponse, error)
	GenerateTempFileName(prefix, suffix string) string
	ParseCarFile(carFilePath string, rootLink *model.RootLink) error
	SliceBigCarFile(carFilePath string) error
	GenerateShardingCarFiles(req *model.CarFileUploadReq, shardingCarFileUploadReqs *[]model.CarFileUploadReq) error
	UploadData(bucketId int, dataPath string) (model.ObjectCreateResponse, error)
	UploadBigCarFile(req *model.CarFileUploadReq) (model.ObjectCreateResponse, error)
	UploadCarFileExt(req *model.CarFileUploadReq, extReader io.Reader) (model.ObjectCreateResponse, error)
	UploadShardingCarFileExt(req *model.CarFileUploadReq, extReader io.Reader) (model.ShardingCarFileUploadResponse, error)
	ImportCarFileExt(req *model.CarFileUploadReq, extReader io.Reader) (model.ObjectCreateResponse, error)
	ImportShardingCarFileExt(req *model.CarFileUploadReq, extReader io.Reader) (model.ShardingCarFileUploadResponse, error)
	ExtractCarFile(carFilePath string, dataDestination string) error
}
