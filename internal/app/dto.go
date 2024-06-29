package app

type MemoryDTO struct {
	Total       uint64
	Free        uint64
	UsedPercent float64
}

type CpuDTO struct {
	UsedPercent float64
	CoreNumber  int
}

type DiskDTO struct {
	Total       uint64
	Free        uint64
	UsedPercent float64
}
