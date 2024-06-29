package monitor

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
)

type Service struct{}

func (s *Service) MemoryUtilization() {
	v, _ := mem.VirtualMemory()

	fmt.Printf("Memory usage: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
}

func (s *Service) CpuUtilization() {
	usage, _ := cpu.Percent(time.Second*1, true)

	for i, cpuUsage := range usage {

		fmt.Printf("CPU core %v Usage: %v,\n", i, cpuUsage)
	}
}

func (s *Service) DiskUtilization() {
	usage, _ := disk.Usage("/")

	fmt.Printf("Disk Usage: %v,\n", usage)
}
