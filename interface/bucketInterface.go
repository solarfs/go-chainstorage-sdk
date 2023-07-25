package _interface

import "github.com/solarfs/go-chainstorage-sdk/model"

type Bucket interface {
	GetBucketList(bucketName string, pageSize, pageIndex int) (model.BucketPageResponse, error)
	CreateBucket(bucketName string, storageNetworkCode, bucketPrincipleCode int) (model.BucketCreateResponse, error)
	EmptyBucket(bucketId int) (model.BucketEmptyResponse, error)
	RemoveBucket(bucketId int, autoEmptyBucketData bool) (model.BucketRemoveResponse, error)
	GetBucketByName(bucketName string) (model.BucketCreateResponse, error)
	GetUsersQuotaByStorageNetworkCode(storageNetworkCode int) (model.UsersQuotaResponse, error)
}
