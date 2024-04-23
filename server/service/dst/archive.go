package dst

import (
	"context"
	"github.com/dstgo/wigfrid/pb"
	"github.com/dstgo/wigfrid/server/handler/dst"
)

func NewArchiveService(archive *dst.ArchiveHandler) *ArchiveService {
	return &ArchiveService{Archive: archive}
}

type ArchiveService struct {
	pb.UnimplementedArchiveServiceServer
	Archive *dst.ArchiveHandler
}

func (a *ArchiveService) Info(ctx context.Context, id *pb.ContainerId) (*pb.ArchiveInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (a *ArchiveService) ListBackups(ctx context.Context, id *pb.ContainerId) (*pb.BackUpList, error) {
	//TODO implement me
	panic("implement me")
}

func (a *ArchiveService) CreateBackup(ctx context.Context, opt *pb.BackupOpt) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (a *ArchiveService) DeleteBackUp(ctx context.Context, opt *pb.BackupOpt) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (a *ArchiveService) RestoreBackUp(ctx context.Context, opt *pb.BackupOpt) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (a *ArchiveService) UploadBackup(ctx context.Context, file *pb.BackupFile) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (a *ArchiveService) DownloadBackup(ctx context.Context, opt *pb.BackupOpt) (*pb.BackupFile, error) {
	//TODO implement me
	panic("implement me")
}

func (a *ArchiveService) mustEmbedUnimplementedArchiveServiceServer() {
	//TODO implement me
	panic("implement me")
}
