package app

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
)

type Service struct{}

func (s *Service) MemoryUtilization() (*MonitoringDTO, error) {
	memoryUsage, err := mem.VirtualMemory()

	return &MonitoringDTO{
		ID:          "MEMORY_01",
		Total:       memoryUsage.Total,
		Free:        memoryUsage.Free,
		UsedPercent: memoryUsage.UsedPercent,
	}, err
}

func (s *Service) CpuUtilization() (MonitoringDTO, error) {
	cpuUsage, err := cpu.Percent(time.Second*3, true)
	var highestCpuUsage MonitoringDTO

	for i, usedPercentage := range cpuUsage {
		if usedPercentage > highestCpuUsage.UsedPercent {
			highestCpuUsage = MonitoringDTO{
				ID:          fmt.Sprintf("CPU_%v", i),
				UsedPercent: usedPercentage,
				Total:       uint64(len(cpuUsage)),
			}
		}
	}

	return highestCpuUsage, err
}

func (s *Service) DiskUtilization() (*MonitoringDTO, error) {
	diskUsage, err := disk.Usage("/")

	return &MonitoringDTO{
		ID:          "DISK_01",
		Total:       diskUsage.Total,
		Free:        diskUsage.Free,
		UsedPercent: diskUsage.UsedPercent,
	}, err
}
