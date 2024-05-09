package types

import (
	"github.com/docker/docker/client"
	"github.com/dstgo/wigfrid/server/conf"
)

type Env struct {
	// docker http api client
	Docker  *client.Client
	AppConf *conf.App
}
