package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"

	"github.com/gorilla/websocket"
)

// Struct to hold system metrics
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

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Upgrade error:", err)
			return
		}
		defer conn.Close()

		for {
			metrics, err := getMetrics()
			if err != nil {
				log.Println("Error getting metrics:", err)
				break
			}

			data, _ := json.Marshal(metrics)
			err = conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println("WriteMessage error:", err)
				break
			}

			time.Sleep(3 * time.Second)
		}
	})

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe error:", err)
	}
}

func getMetrics() (Metrics, error) {
	// cpuPercents, err := cpu.Percent(500*time.Millisecond, false)
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
