package server

import (
	"context"
	"github.com/dstgo/wigfrid/server/conf"
	"github.com/dstgo/wigfrid/server/types"
	"github.com/go-kratos/kratos/v2"
	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"google.golang.org/grpc"
	"log/slog"
)

// NewApp returns new grpc app server
func NewApp(ctx context.Context, appConf *conf.App) (*kratos.App, error) {
	slog.Info("server is initializing")

	// initialize gRPC server
	server := kgrpc.NewServer(
		kgrpc.Network("tcp"),
		kgrpc.Address(appConf.GRPC.Address),
		kgrpc.Options(
			grpc.MaxRecvMsgSize(appConf.GRPC.MaxRecv),
			grpc.MaxSendMsgSize(appConf.GRPC.MaxSend),
			grpc.ReadBufferSize(appConf.GRPC.ReadBuffer),
			grpc.WriteBufferSize(appConf.GRPC.WriteBuffer),
			grpc.MaxHeaderListSize(appConf.GRPC.MaxHeaderSize),
		),
	)

	// load docker client
	slog.Debug("load docker client from local")
	dockerClient, err := loadDockerFromLocal(ctx)
	if err != nil {
		return nil, err
	}

	services := setup(&types.Env{
		Docker:  dockerClient,
		AppConf: appConf,
	})

	// register services
	registerService(server, services)

	app := kratos.New(
		kratos.Context(ctx),
		kratos.Server(server),
		kratos.Logger(&SlogAdapter{slog.Default()}),
	)
	return app, nil
}
