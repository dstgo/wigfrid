package dst

import (
	"context"
	"github.com/dstgo/wigfrid/pb"
	"github.com/dstgo/wigfrid/server/handler/dst"
)

func NewModService(mod *dst.ModHandler) *ModService {
	return &ModService{Mod: mod}
}

type ModService struct {
	pb.UnimplementedModServiceServer
	Mod *dst.ModHandler
}

func (m *ModService) GetWorkShopModList(ctx context.Context, req *pb.ModListReq) (*pb.ModListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (m *ModService) Subscribe(ctx context.Context, id *pb.ModId) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (m *ModService) Unsubscribe(ctx context.Context, id *pb.ModId) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (m *ModService) UpdateMod(ctx context.Context, id *pb.ModId) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (m *ModService) CheckUpdate(ctx context.Context, id *pb.ContainerId) (*pb.CheckUpdateResult, error) {
	//TODO implement me
	panic("implement me")
}

func (m *ModService) GetModSettings(ctx context.Context, id *pb.ContainerId) (*pb.ModSettings, error) {
	//TODO implement me
	panic("implement me")
}

func (m *ModService) SaveModSettings(ctx context.Context, req *pb.SaveModSettingsReq) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (m *ModService) GetRawModSettings(ctx context.Context, id *pb.ContainerId) (*pb.RawModSettings, error) {
	//TODO implement me
	panic("implement me")
}

func (m *ModService) SaveRawModSettings(ctx context.Context, req *pb.SaveRawModSettingsReq) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (m *ModService) mustEmbedUnimplementedModServiceServer() {
	//TODO implement me
	panic("implement me")
}
