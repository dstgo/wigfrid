package daemon

import (
	"context"
	"github.com/docker/docker/client"
	"github.com/dstgo/wigfrid/pb"
	"github.com/dstgo/wigfrid/pkg/hostinfo"
	"github.com/dstgo/wigfrid/server/conf"
	"github.com/shirou/gopsutil/v3/host"
	"runtime"
)

func NewDaemonHandler(client *client.Client, conf *conf.App) *DaemonHandler {
	return &DaemonHandler{docker: client, conf: conf}
}

type DaemonHandler struct {
	docker *client.Client
	conf   *conf.App
}

// HostInfo returns basic information about local machine
func (d *DaemonHandler) HostInfo(ctx context.Context) (*pb.SystemInfo, error) {
	hostStat, err := host.Info()
	if err != nil {
		return nil, err
	}

	// docker information
	dockerStat, err := d.docker.Info(ctx)
	if err != nil {
		return nil, err
	}

	dockerInfo := &pb.DockerInfo{
		Containers:    int64(dockerStat.Containers),
		Running:       int64(dockerStat.ContainersRunning),
		Pause:         int64(dockerStat.ContainersPaused),
		Stopped:       int64(dockerStat.ContainersStopped),
		Images:        int64(dockerStat.Images),
		Driver:        dockerStat.Driver,
		Version:       dockerStat.ServerVersion,
		KernelVersion: dockerStat.KernelVersion,
	}

	// cpu information
	cpuStat, err := hostinfo.CpuStat()
	if err != nil {
		return nil, err
	}

	// memory information
	memStat, err := hostinfo.MemStat()
	if err != nil {
		return nil, err
	}

	// root fs information
	rootFsStat, err := hostinfo.RootFsStat()
	if err != nil {
		return nil, err
	}

	// basic host information
	sysInfo := &pb.SystemInfo{
		Os:           hostStat.OS,
		Arch:         hostStat.KernelArch,
		GoVersion:    runtime.Version(),
		BuildVersion: d.conf.Version,
		Docker:       dockerInfo,
		CpuModel:     cpuStat.Model,
		Resource: &pb.Resource{
			Cpu:    cpuStat.Count,
			Memory: memStat.Total,
			Disk:   int64(rootFsStat.Total),
		},
	}

	return sysInfo, nil
}

// HealthCheck returns the cpu usage and memory usage information immediately without sleeping.
func (d *DaemonHandler) HealthCheck(ctx context.Context) (*pb.HealthInfo, error) {

	// cpu information
	cpuStat, err := hostinfo.CpuStat()
	if err != nil {
		return nil, err
	}

	// cpu usage
	usage, err := hostinfo.CpuUsage()
	if err != nil {
		return nil, err
	}

	// memory information
	memStat, err := hostinfo.MemStat()
	if err != nil {
		return nil, err
	}

	healthInfo := &pb.HealthInfo{
		Cpu: &pb.CpuHealth{
			Count: cpuStat.Count,
			Usage: usage,
		},
		Mem: &pb.MemoryHealth{
			Total: memStat.Total,
			Used:  memStat.Used,
			Free:  memStat.Free,
			Usage: memStat.Percent,
		},
	}

	return healthInfo, nil
}
