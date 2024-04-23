package server

import (
	"context"
	"github.com/dstgo/wigfrid/server/conf"
	"github.com/go-kratos/kratos/v2"
	"log/slog"
)

// NewApp returns new grpc app server
func NewApp(ctx context.Context, appConf *conf.App) (*kratos.App, error) {
	slog.Info("server is initializing")
	app := kratos.New()
	return app, nil
}
