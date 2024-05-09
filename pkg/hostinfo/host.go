package hostinfo

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

func init() {
	cpu.Percent(-1, false)
}

type CpuInfo struct {
	Count int64
	cpu.InfoStat
}

// CpuStat returns host cpu information
func CpuStat() (CpuInfo, error) {
	var cpuInfo CpuInfo
	stats, err := cpu.Info()
	if err != nil {
		return cpuInfo, err
	}

	if len(stats) > 1 {
		cpuInfo.InfoStat = stats[0]
	}

	counts, err := cpu.Counts(true)
	if err != nil {
		return cpuInfo, err
	}
	cpuInfo.Count = int64(counts)

	return cpuInfo, nil
}

// CpuUsage returns cpu usage percent from last call
func CpuUsage() (float64, error) {
	percent, err := cpu.Percent(-1, false)
	if err != nil {
		return 0, err
	}
	return percent[0], nil
}

type MemInfo struct {
	Total   int64
	Used    int64
	Free    int64
	Percent float64
}

// MemStat returns host mem information
func MemStat() (MemInfo, error) {
	memory, err := mem.SwapMemory()
	if err != nil {
		return MemInfo{}, err
	}
	return MemInfo{
		Total:   int64(memory.Total),
		Used:    int64(memory.Used),
		Free:    int64(memory.Free),
		Percent: memory.UsedPercent,
	}, nil
}

// RootFsStat return rootfs info of host
func RootFsStat() (*disk.UsageStat, error) {
	usage, err := disk.Usage("/")
	if err != nil {
		return nil, err
	}
	return usage, nil
}
