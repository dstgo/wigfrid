package server

import (
	"github.com/dstgo/wigfrid/assets"
	"github.com/dstgo/wigfrid/pb"
	"github.com/dstgo/wigfrid/server/conf"
	"github.com/dstgo/wigfrid/server/service"
	"github.com/dstgo/wigfrid/server/types"
	"github.com/ginx-contribs/logx"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"io"
)

var EnvProvider = wire.NewSet(
	wire.FieldsOf(new(*types.Env), "Docker"),
)

// PrintBanner writes the banner into given writer
func PrintBanner(writer io.Writer) error {
	bytes, err := assets.Fs.ReadFile("banner.txt")
	if err != nil {
		return err
	}
	_, err = writer.Write(bytes)
	return err
}

// NewLogger returns a new app logger with the given options
func NewLogger(option conf.Log) (*logx.Logger, error) {

	writer, err := logx.NewWriter(&logx.WriterOptions{
		Filename: option.Filename,
	})
	if err != nil {
		return nil, err
	}
	handler, err := logx.NewHandler(writer, &logx.HandlerOptions{
		Level:       option.Level,
		Format:      option.Format,
		Prompt:      option.Prompt,
		Source:      option.Source,
		ReplaceAttr: nil,
		Color:       option.Color,
	})
	if err != nil {
		return nil, err
	}
	logger, err := logx.New(
		logx.WithHandlers(handler),
	)
	if err != nil {
		return nil, err
	}
	return logger, nil
}

// register grpc router
func registerService(server *grpc.Server, service service.Service) {
	// cronjob
	pb.RegisterCronJobServiceServer(server, service.CronJob)
	// daemon
	pb.RegisterDaemonServiceServer(server, service.Daemon)
	pb.RegisterContainerServiceServer(server, service.Container)
	pb.RegisterImageServiceServer(server, service.Image)
	// dst
	pb.RegisterArchiveServiceServer(server, service.Archive)
	pb.RegisterModServiceServer(server, service.Mod)
	pb.RegisterPlayerServiceServer(server, service.Player)
	pb.RegisterShardServiceServer(server, service.Shard)
	pb.RegisterSettingServiceServer(server, service.Setting)
}
