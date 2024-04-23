package cronjob

import (
	"context"
	"github.com/dstgo/wigfrid/pb"
	"github.com/dstgo/wigfrid/server/handler/cronjob"
)

func NewCronJobService(cronJob *cronjob.CronJobHandler) *CronJobService {
	return &CronJobService{CronJob: cronJob}
}

type CronJobService struct {
	pb.UnimplementedCronJobServiceServer
	CronJob *cronjob.CronJobHandler
}

func (c *CronJobService) Create(ctx context.Context, req *pb.CreateJobReq) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CronJobService) Delete(ctx context.Context, req *pb.CreateJobReq) (*pb.Notify, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CronJobService) List(ctx context.Context, req *pb.CreateJobReq) (*pb.JobList, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CronJobService) mustEmbedUnimplementedCronJobServiceServer() {
	//TODO implement me
	panic("implement me")
}
