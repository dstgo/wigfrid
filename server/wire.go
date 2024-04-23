//go:build wireinject

// The build tag makes sure the stub is not built in the final build
package server

import (
	"github.com/dstgo/wigfrid/server/handler"
	"github.com/dstgo/wigfrid/server/service"
	"github.com/dstgo/wigfrid/server/types"
	"github.com/google/wire"
)

func Setup(env *types.Env) service.Service {
	panic(wire.Build(
		EnvProvider,
		handler.Provider,
		service.Provider,
	))
}
