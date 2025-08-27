package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
)

type Metrics struct {
	CPUPercent float64 `json:"cpu_percent"`
	Memory     uint64  `json:"memory_total"`
	MemoryUsed uint64  `json:"memory_used"`
	MemoryFree uint64  `json:"memory_free"`
	DiskTotal  uint64  `json:"disk_total"`
	DiskUsed   uint64  `json:"disk_used"`
	DiskFree   uint64  `json:"disk_free"`
	NetSent    uint64  `json:"net_sent"`
	NetRecv    uint64  `json:"net_recv"`
}

func main() {
	metrics, err := getMetrics()
	if err != nil {
		panic(err)
	}
	// Print the metrics to the console (or handle them as needed)
	fmt.Printf("Metrics: %+v\n", metrics)
}

func getMetrics() (Metrics, error) {
	cpuPercents, err := cpu.Percent(0, false)
	if err != nil {
		return Metrics{}, err
	}
	cpuPercent := cpuPercents[0]

	v, err := mem.VirtualMemory()
	if err != nil {
		return Metrics{}, err
	}

	d, err := disk.Usage("/")
	if err != nil {
		return Metrics{}, err
	}

	netIOs, err := net.IOCounters(false)
	if err != nil {
		return Metrics{}, err
	}
	netIO := netIOs[0]

	metrics := Metrics{
		CPUPercent: cpuPercent,
		Memory:     v.Total,
		MemoryUsed: v.Used,
		MemoryFree: v.Free,
		DiskTotal:  d.Total,
		DiskUsed:   d.Used,
		DiskFree:   d.Free,
		NetSent:    netIO.BytesSent,
		NetRecv:    netIO.BytesRecv,
	}

	return metrics, nil
}
