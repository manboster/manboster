package system

import (
	"context"
	"fmt"
	"runtime"
	"strings"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/load"
	"github.com/shirou/gopsutil/v4/mem"
)

type SysInfo struct {
	System      string     `json:"system"`
	Arch        string     `json:"arch"`
	Hostname    string     `json:"hostname"`
	KernelVer   string     `json:"kernel_ver"`
	Uptime      uint64     `json:"uptime"`
	CPUs        []CPUInfo  `json:"cpus"`
	MemoryTotal uint64     `json:"memory_total"`
	MemoryAvail uint64     `json:"memory_avail"`
	DiskTotal   uint64     `json:"disk_total"`
	DiskAvail   uint64     `json:"disk_avail"`
	LoadAvg     [3]float64 `json:"load_avg"`
}

type CPUInfo struct {
	Name      string  `json:"name"`
	Vendor    string  `json:"vendor"`
	Family    string  `json:"family"`
	Model     string  `json:"model"`
	Cores     int32   `json:"cores"`
	MHz       float64 `json:"mhz"`
	CacheSize int32   `json:"cache_kb"`
}

func getSystemInfo(ctx context.Context) (*SysInfo, error) {
	info := &SysInfo{
		System: runtime.GOOS,
		Arch:   runtime.GOARCH,
	}

	h, err := host.InfoWithContext(ctx)
	if err != nil {
		return nil, err
	}
	info.Hostname = h.Hostname
	info.KernelVer = h.KernelVersion
	info.Uptime = h.Uptime

	cpuStats, err := cpu.InfoWithContext(ctx)
	if err != nil {
		return nil, err
	}
	for _, c := range cpuStats {
		info.CPUs = append(info.CPUs, CPUInfo{
			Name:      c.ModelName,
			Vendor:    c.VendorID,
			Family:    c.Family,
			Model:     c.Model,
			Cores:     c.Cores,
			MHz:       c.Mhz,
			CacheSize: c.CacheSize,
		})
	}

	m, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		return nil, err
	}
	info.MemoryTotal = m.Total
	info.MemoryAvail = m.Available

	total, avail, err := getDiskUsage(ctx)
	if err != nil {
		return nil, err
	}
	info.DiskTotal = total
	info.DiskAvail = avail

	l, err := load.AvgWithContext(ctx)
	if err != nil {
		return nil, err
	}
	info.LoadAvg = [3]float64{l.Load1, l.Load5, l.Load15}

	return info, nil
}

func getDiskUsage(ctx context.Context) (total, avail uint64, err error) {
	parts, err := disk.PartitionsWithContext(ctx, false)
	if err != nil {
		return 0, 0, err
	}
	for _, p := range parts {
		if isSystemPartition(p.Mountpoint) {
			usage, err := disk.UsageWithContext(ctx, p.Mountpoint)
			if err != nil {
				return 0, 0, err
			}
			return usage.Total, usage.Free, nil
		}
	}
	return 0, 0, fmt.Errorf("no system partition found")
}

func isSystemPartition(mount string) bool {
	if runtime.GOOS == "windows" {
		return strings.HasPrefix(strings.ToUpper(mount), "C:")
	}
	return mount == "/"
}
