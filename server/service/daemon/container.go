package daemon

import (
	context "context"
	"github.com/dstgo/wigfrid/pb"
	"github.com/dstgo/wigfrid/server/handler/daemon"
)

func NewContainerService(container *daemon.ContainerHandler) *ContainerService {
	return &ContainerService{Container: container}
}

type ContainerService struct {
	Container *daemon.ContainerHandler
	pb.UnimplementedContainerServiceServer
}

func (c *ContainerService) List(ctx context.Context, req *pb.ListContainerReq) (*pb.ListContainerResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ContainerService) Create(ctx context.Context, req *pb.CreateContainerReq) (*pb.CreateContainerResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ContainerService) UpdateQuota(ctx context.Context, resource *pb.Resource) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ContainerService) Log(ctx context.Context, req *pb.LogContainerReq) (*pb.ContainerLog, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ContainerService) Stats(ctx context.Context, id *pb.ContainerId) (*pb.HealthInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ContainerService) Start(ctx context.Context, id *pb.ContainerId) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ContainerService) Stop(ctx context.Context, id *pb.ContainerId) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ContainerService) Restart(ctx context.Context, id *pb.ContainerId) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ContainerService) Delete(ctx context.Context, id *pb.ContainerId) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ContainerService) ForceDelete(ctx context.Context, id *pb.ContainerId) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ContainerService) mustEmbedUnimplementedContainerServiceServer() {
	//TODO implement me
	panic("implement me")
}
