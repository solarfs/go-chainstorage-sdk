package _interface

import "github.com/solarfs/go-chainstorage-sdk/model"

type Object interface {
	GetObjectList(bucketId int, objectItem string, pageSize, pageIndex int) (model.ObjectPageResponse, error)
	RemoveObject(objectIds []int) (model.ObjectRemoveResponse, error)
	RenameObject(objectId int, objectName string, isOverwrite bool) (model.ObjectRenameResponse, error)
	MarkObject(objectId int, isMarked bool) (model.ObjectMarkResponse, error)
	IsExistObjectByCid(objectCid string) (model.ObjectExistResponse, error)
	GetObjectByName(bucketId int, objectName string) (model.ObjectCreateResponse, error)
}
