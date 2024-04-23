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

func (d DaemonService) HostInfo(ctx context.Context, empty *emptypb.Empty) (*pb.SystemInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (d DaemonService) HealthCheck(ctx context.Context, empty *emptypb.Empty) (*pb.HealthInfo, error) {
	//TODO implement me
	panic("implement me")
}
