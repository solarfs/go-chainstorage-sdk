package model

import "time"

type Gateway struct {
	Id                 int       `json:"id" comment:"网关ID"`
	UserId             int       `json:"-" comment:"用户ID"`
	IsBindBucket       bool      `json:"isBindBucket" comment:"记录绑定bucket状态（true-绑定，false-未绑定）,不能和IsBindObject同时为true"`
	BucketId           int       `json:"bucketId" comment:"桶ID"`
	BucketName         string    `json:"bucketName" comment:"桶名称（3-63字长度限制）"`
	IsBindObject       bool      `json:"isBindObject" comment:"记录绑定Object状态（true-绑定，false-未绑定）,不能和IsBindBucket同时为true"`
	ObjectId           int64     `json:"objectId" comment:"ObjectID"`
	StorageNetworkCode int       `json:"storageNetworkCode" comment:"存储网络编码（10001-IPFS）"`
	GatewayName        string    `json:"gatewayName" comment:"Gateway名称,默认的域名前缀（1-63字长度限制,只支持数字小写字符中横线）"`
	IngressName        string    `json:"ingressName" comment:"IngressName,默认{GatewayName}-gw（1-100字长度限制）"`
	IngressHost        string    `json:"ingressHost" comment:"IngressHost,访问的域名（1-100字长度限制）"`
	KongRouteID        string    `json:"kongRouteID" comment:"KongRouteID（默认36字节）"`
	KongRouteName      string    `json:"kongRouteName" comment:"KongRouteName（1-256字长度限制）"`
	UsedNetworkTraffic int64     `json:"usedNetworkTraffic" comment:"本月使用流量（字节）暂时未持久化,查询直接取prometheus"`
	UsedRequestCount   int64     `json:"usedRequestCount" comment:"本月请求数量（次）暂时未持久化,查询直接取prometheus"`
	Object             *Object   `json:"object" comment:"对象"`
	CreatedAt          time.Time `json:"createdAt" comment:"创建时间"`
	UpdatedAt          time.Time `json:"updatedAt" comment:"最后更新时间"`
}
