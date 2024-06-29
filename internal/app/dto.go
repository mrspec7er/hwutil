package app

type MonitoringDTO struct {
	ID          string
	Total       uint64
	Free        uint64
	UsedPercent float64
}
