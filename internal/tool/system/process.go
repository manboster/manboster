package system

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/shirou/gopsutil/v4/process"
)

type ProcessInfo struct {
	PID    int32   `json:"pid"`
	Name   string  `json:"name"`
	CPU    float64 `json:"cpu_percent"`
	Mem    float32 `json:"mem_percent"`
	Status string  `json:"status"`
}

func listProcesses(ctx context.Context) ([]ProcessInfo, error) {
	procs, err := process.ProcessesWithContext(ctx)
	if err != nil {
		return nil, err
	}

	var list []ProcessInfo
	for _, p := range procs {
		name, _ := p.NameWithContext(ctx)
		cpuPercent, _ := p.CPUPercentWithContext(ctx)
		memPercent, _ := p.MemoryPercentWithContext(ctx)
		status, _ := p.StatusWithContext(ctx)
		if name == "" {
			continue
		}

		list = append(list, ProcessInfo{
			PID:    p.Pid,
			Name:   name,
			CPU:    cpuPercent,
			Mem:    memPercent,
			Status: strings.Join(status, ","),
		})
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].CPU > list[j].CPU
	})

	return list, nil
}

func getProcessInfo(ctx context.Context, pid int32) (*ProcessInfo, error) {
	p, err := process.NewProcessWithContext(ctx, pid)
	if err != nil {
		return nil, fmt.Errorf("process %d not found: %w", pid, err)
	}

	name, _ := p.NameWithContext(ctx)
	cpuPercent, _ := p.CPUPercentWithContext(ctx)
	memPercent, _ := p.MemoryPercentWithContext(ctx)
	status, _ := p.StatusWithContext(ctx)

	return &ProcessInfo{
		PID:    pid,
		Name:   name,
		CPU:    cpuPercent,
		Mem:    memPercent,
		Status: strings.Join(status, ","),
	}, nil
}

func killProcess(ctx context.Context, pid int32) error {
	p, err := process.NewProcessWithContext(ctx, pid)
	if err != nil {
		return fmt.Errorf("process %d not found: %w", pid, err)
	}
	return p.KillWithContext(ctx)
}
