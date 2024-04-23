// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package server

import (
	"github.com/dstgo/wigfrid/server/handler/cronjob"
	"github.com/dstgo/wigfrid/server/handler/daemon"
	"github.com/dstgo/wigfrid/server/handler/dst"
	"github.com/dstgo/wigfrid/server/service"
	cronjob2 "github.com/dstgo/wigfrid/server/service/cronjob"
	daemon2 "github.com/dstgo/wigfrid/server/service/daemon"
	dst2 "github.com/dstgo/wigfrid/server/service/dst"
	"github.com/dstgo/wigfrid/server/types"
)

// Injectors from wire.go:

func Setup(env *types.Env) service.Service {
	cronJobHandler := cronjob.NewCronJobHandler()
	cronJobService := cronjob2.NewCronJobService(cronJobHandler)
	client := env.Docker
	daemonHandler := daemon.NewDaemonHandler(client)
	daemonService := daemon2.NewDaemonService(daemonHandler)
	containerHandler := daemon.NewContainerHandler(client)
	containerService := daemon2.NewContainerService(containerHandler)
	imageHandler := daemon.NewImageHandler(client)
	imageService := daemon2.NewImageService(imageHandler)
	archiveHandler := dst.NewArchiveHandler()
	archiveService := dst2.NewArchiveService(archiveHandler)
	modHandler := dst.NewModHandler()
	modService := dst2.NewModService(modHandler)
	playerHandler := dst.NewPlayerHandler()
	playerService := dst2.NewPlayerService(playerHandler)
	shardHandler := dst.NewShardHandler()
	shardService := dst2.NewShardService(shardHandler)
	settingHandler := dst.NewSettingHandler()
	settingService := dst2.NewSettingService(settingHandler)
	serviceService := service.Service{
		CronJob:   cronJobService,
		Daemon:    daemonService,
		Container: containerService,
		Image:     imageService,
		Archive:   archiveService,
		Mod:       modService,
		Player:    playerService,
		Shard:     shardService,
		Setting:   settingService,
	}
	return serviceService
}
