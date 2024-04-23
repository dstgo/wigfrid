package daemon

import (
	"context"
	"github.com/dstgo/wigfrid/pb"
	"github.com/dstgo/wigfrid/server/handler/daemon"
)

func NewImageService(image *daemon.ImageHandler) *ImageService {
	return &ImageService{Image: image}
}

type ImageService struct {
	pb.UnimplementedImageServiceServer
	Image *daemon.ImageHandler
}

func (i ImageService) List(ctx context.Context, req *pb.ListImagesReq) (*pb.ListImageResult, error) {
	//TODO implement me
	panic("implement me")
}

func (i ImageService) Remove(ctx context.Context, id *pb.ImageId) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}
