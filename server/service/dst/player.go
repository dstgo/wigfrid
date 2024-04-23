package dst

import (
	"context"
	"github.com/dstgo/wigfrid/pb"
	"github.com/dstgo/wigfrid/server/handler/dst"
)

func NewPlayerService(player *dst.PlayerHandler) *PlayerService {
	return &PlayerService{Player: player}
}

type PlayerService struct {
	Player *dst.PlayerHandler
	pb.UnimplementedPlayerServiceServer
}

func (p *PlayerService) GetPlayerStats(ctx context.Context, id *pb.ContainerId) (*pb.PlayerStatisticInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PlayerService) GetPlayerChatLog(ctx context.Context, id *pb.ContainerId) (*pb.PlayerChatLog, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PlayerService) ExecutePlayer(ctx context.Context, req *pb.ExecutePlayerReq) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PlayerService) GetWhiteList(ctx context.Context, id *pb.ContainerId) (*pb.PlayerListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PlayerService) GetBlackList(ctx context.Context, id *pb.ContainerId) (*pb.PlayerListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PlayerService) GetAdminList(ctx context.Context, id *pb.ContainerId) (*pb.PlayerListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PlayerService) AddWhiteList(ctx context.Context, req *pb.PlayerListReq) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PlayerService) AddBlackList(ctx context.Context, req *pb.PlayerListReq) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PlayerService) AddAdminList(ctx context.Context, req *pb.PlayerListReq) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PlayerService) RemoveWhiteList(ctx context.Context, req *pb.PlayerListReq) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PlayerService) RemoveBlackList(ctx context.Context, req *pb.PlayerListReq) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PlayerService) RemoveAdminList(ctx context.Context, req *pb.PlayerListReq) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PlayerService) mustEmbedUnimplementedPlayerServiceServer() {
	//TODO implement me
	panic("implement me")
}
