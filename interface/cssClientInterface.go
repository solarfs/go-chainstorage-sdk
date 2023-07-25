package _interface

import "github.com/solarfs/go-chainstorage-sdk/model"

type CssClient interface {
	GetIpfsVersion() (model.VersionResponse, error)
	GetApiVersion() (model.VersionResponse, error)
}
