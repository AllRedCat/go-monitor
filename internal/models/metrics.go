package models

type Metrics struct {
	CPUPercent  float64 `json:"cpu_percent"`
	Memory      uint64  `json:"memory_total"`
	MemoryUsed  uint64  `json:"memory_used"`
	MemoryFree  uint64  `json:"memory_free"`
	DiskTotal   uint64  `json:"disk_total"`
	DiskUsed    uint64  `json:"disk_used"`
	DiskFree    uint64  `json:"disk_free"`
	DiskPercent float64 `json:"disk_percent"`
	NetSent     uint64  `json:"net_sent"`
	NetRecv     uint64  `json:"net_recv"`
}