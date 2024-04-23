package dst

import (
	"context"
	"github.com/dstgo/wigfrid/pb"
	"github.com/dstgo/wigfrid/server/handler/dst"
)

func NewShardService(shards *dst.ShardHandler) *ShardService {
	return &ShardService{Shards: shards}
}

type ShardService struct {
	Shards *dst.ShardHandler
	pb.UnimplementedShardServiceServer
}

func (s *ShardService) Start(ctx context.Context, request *pb.ControlRequest) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ShardService) ReStart(ctx context.Context, request *pb.ControlRequest) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ShardService) Stop(ctx context.Context, request *pb.ControlRequest) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ShardService) State(ctx context.Context, request *pb.ControlRequest) (*pb.StateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ShardService) Logs(ctx context.Context, request *pb.LogsRequest) (*pb.LogsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ShardService) Execute(ctx context.Context, request *pb.CommandRequest) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ShardService) Version(ctx context.Context, id *pb.ContainerId) (*pb.VersionResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ShardService) Update(ctx context.Context, id *pb.ContainerId) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ShardService) RollBack(ctx context.Context, id *pb.ContainerId) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ShardService) Reset(ctx context.Context, id *pb.ContainerId) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ShardService) Clear(ctx context.Context, id *pb.ContainerId) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ShardService) mustEmbedUnimplementedShardServiceServer() {
	//TODO implement me
	panic("implement me")
}
