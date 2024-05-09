package daemon

import (
	"context"
	"github.com/dstgo/wigfrid/pb"
	"github.com/dstgo/wigfrid/server/handler/daemon"
	"google.golang.org/protobuf/types/known/emptypb"
)

func NewDaemonService(daemon *daemon.DaemonHandler) *DaemonService {
	return &DaemonService{Daemon: daemon}
}

type DaemonService struct {
	pb.UnimplementedDaemonServiceServer
	Daemon *daemon.DaemonHandler
}

// HostInfo returns basic information about local machine
func (d DaemonService) HostInfo(ctx context.Context, empty *emptypb.Empty) (*pb.SystemInfo, error) {
	return d.Daemon.HostInfo(ctx)
}

func (d DaemonService) HealthCheck(ctx context.Context, empty *emptypb.Empty) (*pb.HealthInfo, error) {
	return d.Daemon.HealthCheck(ctx)
}
