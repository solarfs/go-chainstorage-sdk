package _interface

import "github.com/paradeum-team/chainstorage-sdk/model"

type CssClient interface {
	GetIpfsVersion() (model.VersionResponse, error)
	GetApiVersion() (model.VersionResponse, error)
}
