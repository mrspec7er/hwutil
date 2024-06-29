package app

import (
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
)

type Service struct{}

func (s *Service) MemoryUtilization() (*MemoryDTO, error) {
	memoryUsage, err := mem.VirtualMemory()

	return &MemoryDTO{
		Total:       memoryUsage.Total,
		Free:        memoryUsage.Free,
		UsedPercent: memoryUsage.UsedPercent,
	}, err
}

func (s *Service) CpuUtilization() ([]*CpuDTO, error) {
	cpuUsage, err := cpu.Percent(time.Second*1, true)
	var cpuUsagePercentage []*CpuDTO

	for i, c := range cpuUsage {
		cpuUsagePercentage = append(cpuUsagePercentage, &CpuDTO{UsedPercent: c, CoreNumber: i})
	}

	return cpuUsagePercentage, err
}

func (s *Service) DiskUtilization() (*DiskDTO, error) {
	diskUsage, err := disk.Usage("/")

	return &DiskDTO{
		Total:       diskUsage.Total,
		Free:        diskUsage.Free,
		UsedPercent: diskUsage.UsedPercent,
	}, err
}
