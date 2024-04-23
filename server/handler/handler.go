package handler

import (
	"github.com/dstgo/wigfrid/server/handler/cronjob"
	"github.com/dstgo/wigfrid/server/handler/daemon"
	"github.com/dstgo/wigfrid/server/handler/dst"
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	// cron job
	cronjob.NewCronJobHandler,
	// daemon
	daemon.NewDaemonHandler,
	daemon.NewImageHandler,
	daemon.NewContainerHandler,
	// dst
	dst.NewArchiveHandler,
	dst.NewModHandler,
	dst.NewPlayerHandler,
	dst.NewSettingHandler,
	dst.NewShardHandler,
)
