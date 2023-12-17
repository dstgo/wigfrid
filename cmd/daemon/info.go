package main

import (
	"github.com/dstgo/size"
	"github.com/dstgo/wigfrid/assets"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/spf13/cobra"
	"os"
	"runtime"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "show the basic information about the wigfird daemon",
	Long:  "show the basic information about the wigfird daemon",
	RunE: func(cmd *cobra.Command, args []string) error {
		return PrintInfo()
	},
}

func PrintInfo() error {
	// os information
	hostInfo, err := host.Info()
	if err != nil {
		return err
	}

	info, err := cpu.Info()
	if err != nil {
		return err
	}

	// cpu information
	cores, err := cpu.Counts(true)
	if err != nil {
		return err
	}

	// memory information
	memory, err := mem.VirtualMemory()
	if err != nil {
		return err
	}

	data := map[string]any{
		"os":        hostInfo.OS,
		"arch":      hostInfo.KernelArch,
		"cpu":       info[0].ModelName,
		"cpuCores":  cores,
		"mem":       size.New(float64(memory.Total), size.B),
		"author":    Author,
		"version":   Version,
		"buildTime": BuildTime,
		"goVersion": runtime.Version(),
	}

	return assets.WriteBanner(os.Stdout, data)
}
