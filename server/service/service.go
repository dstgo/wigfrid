package service

import (
	"github.com/dstgo/wigfrid/server/service/cronjob"
	"github.com/dstgo/wigfrid/server/service/daemon"
	"github.com/dstgo/wigfrid/server/service/dst"
	"github.com/google/wire"
)

type Service struct {
	CronJob   *cronjob.CronJobService
	Daemon    *daemon.DaemonService
	Container *daemon.ContainerService
	Image     *daemon.ImageService
	Archive   *dst.ArchiveService
	Mod       *dst.ModService
	Player    *dst.PlayerService
	Shard     *dst.ShardService
	Setting   *dst.SettingService
}

var Provider = wire.NewSet(
	// cronjob
	cronjob.NewCronJobService,
	// daemon
	daemon.NewDaemonService,
	daemon.NewContainerService,
	daemon.NewImageService,
	// dst
	dst.NewArchiveService,
	dst.NewModService,
	dst.NewPlayerService,
	dst.NewShardService,
	dst.NewSettingService,
	// service
	wire.Struct(new(Service), "*"),
)
