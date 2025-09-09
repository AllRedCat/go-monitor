package services

import (
	"time"

	"monitor/internal/models"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
)

func GetMetrics() (models.Metrics, error) {
	cpuPercents, err := cpu.Percent(1000*time.Millisecond, false)
	if err != nil {
		return models.Metrics{}, err
	}
	cpuPercent := cpuPercents[0]

	v, err := mem.VirtualMemory()
	if err != nil {
		return models.Metrics{}, err
	}

	d, err := disk.Usage("/")
	if err != nil {
		return models.Metrics{}, err
	}

	netIOs, err := net.IOCounters(false)
	if err != nil {
		return models.Metrics{}, err
	}
	netIO := netIOs[0]

	metrics := models.Metrics{
		CPUPercent:  cpuPercent,
		Memory:      v.Total,
		MemoryUsed:  v.Used,
		MemoryFree:  v.Free,
		DiskTotal:   d.Total,
		DiskUsed:    d.Used,
		DiskFree:    d.Free,
		DiskPercent: d.UsedPercent,
		NetSent:     netIO.BytesSent,
		NetRecv:     netIO.BytesRecv,
	}

	return metrics, nil
}
