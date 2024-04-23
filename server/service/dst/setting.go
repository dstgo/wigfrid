package dst

import (
	"context"
	"github.com/dstgo/wigfrid/pb"
	"github.com/dstgo/wigfrid/server/handler/dst"
)

func NewSettingService(setting *dst.SettingHandler) *SettingService {
	return &SettingService{Setting: setting}
}

type SettingService struct {
	Setting *dst.SettingHandler
	pb.UnimplementedSettingServiceServer
}

func (s *SettingService) GetRoomSetting(ctx context.Context, id *pb.ContainerId) (*pb.RoomSetting, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SettingService) SaveRoomSetting(ctx context.Context, setting *pb.RoomSetting) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SettingService) GetWorldSetting(ctx context.Context, id *pb.ContainerId) (*pb.WorldSetting, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SettingService) SaveWorldSetting(ctx context.Context, setting *pb.WorldSetting) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SettingService) GetRawWorldSetting(ctx context.Context, id *pb.ContainerId) (*pb.RawWorldSetting, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SettingService) SaveRawWorldSetting(ctx context.Context, setting *pb.RawWorldSetting) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SettingService) mustEmbedUnimplementedSettingServiceServer() {
	//TODO implement me
	panic("implement me")
}
