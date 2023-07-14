package _interface

import "github.com/solafs/go-chainstorage-sdk/model"

type CssClient interface {
	GetIpfsVersion() (model.VersionResponse, error)
	GetApiVersion() (model.VersionResponse, error)
}
