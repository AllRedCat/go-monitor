package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/docker"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"

	"github.com/gorilla/websocket"
)

// Struct to hold system metrics
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

type CPUUsage struct {
	GeneralPercent float64   `json:"general_percent"`
	EachPercent    []float64 `json:"each_percent"`
}

type CPUInfo struct {
	CPUI        int32   `json:"cpu"`
	SteppingI   int32   `json:"stepping"`
	PhysicalIdI string  `json:"physicalId`
	CoreIdI     string  `json:"coreId"`
	CoresI      int32   `json:"cores"`
	ModelNameI  string  `json:"modelName"`
	MhzI        float64 `json:"mhz"`
	ChaceSizeI  int32   `json:"cacheSize"`
	MicrocodeI  string  `json:"microencode"`
}

// Structure for the request body in "/exec"
type RequestPath struct {
	Path string `json:"path"`
}

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	// Route to get machine info
	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		// Check if the Method it's GET
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		info, err := getMachineInfo() // Call function to get system info
		// Check if an error appears
		if err != nil {
			log.Println("Error getting info:", err)
			json.NewEncoder(w).Encode(err)
		}

		// Send info in JSON format
		json.NewEncoder(w).Encode(info)
	})

	http.HandleFunc("/docker", func(w http.ResponseWriter, r *http.Request) {
		// Check if the Method it's GET
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		doc, err := getContainers() // Cal function to get containers info
		// Check if an error appears
		if err != nil {
			log.Println("Error getting dockers:", err)
			json.NewEncoder(w).Encode(err)
			return
		}

		// If there are no containers, returns success with code 204 (no content)
		if len(doc) == 0 {
			// http.Error(w, "Not found", http.StatusNotFound)
			w.WriteHeader(204)
			return
		}

		// Return a JSON with containers info
		json.NewEncoder(w).Encode(doc)
	})

	http.HandleFunc("/exec", func(w http.ResponseWriter, r *http.Request) {
		// Check if the Method it's POST
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Check if there is a body
		if r.Body == http.NoBody {
			http.Error(w, "A body is required", http.StatusBadRequest)
			return
		}

		// Variable to allocate the path recived in the request body
		var req RequestPath
		// Verify that the body has been successfully translated to JSON
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Verify if body have "scriptPath"
		if req.Path == "" {
			http.Error(w, `"path" is required`, http.StatusBadRequest)
			return
		}

		// Exec the script
		cmd := exec.Command("bash", req.Path)
		output, err := cmd.CombinedOutput()

		// Build a response
		resp := map[string]interface{}{
			"output": string(output),
			"error":  nil,
		}

		// Get a possible error
		if err != nil {
			resp["error"] = err.Error()
		}

		// Set the response headers
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	http.HandleFunc("/ws/cpu", func(w http.ResponseWriter, r *http.Request) {
		connect, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Error: ", err)
			return
		}
		defer connect.Close()

		for {
			cpuUsage, err := getCpuUsage()
			if err != nil {
				log.Println("Error getting usage:", err)
				break
			}

			data, _ := json.Marshal(cpuUsage)
			err = connect.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println("WriterMessage error:", err)
				break
			}

			time.Sleep(3 * time.Second)
		}
	})

	// WebSocket to get metrics with a time of 3 seconds
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

func getMachineInfo() ([]CPUInfo, error) {
	cpus, err := cpu.Info()
	if err != nil {
		return nil, err
	}

	var infos []CPUInfo
	for _, i := range cpus {
		infos = append(infos, CPUInfo{
			CPUI:        i.CPU,
			SteppingI:   i.Stepping,
			PhysicalIdI: i.PhysicalID,
			CoreIdI:     i.CoreID,
			CoresI:      i.Cores,
			ModelNameI:  i.ModelName,
			MhzI:        i.Mhz,
			ChaceSizeI:  i.CacheSize,
			MicrocodeI:  i.Microcode,
		})
	}

	return infos, nil
}

func getMetrics() (Metrics, error) {
	cpuPercents, err := cpu.Percent(1000*time.Millisecond, false)
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

func getCpuUsage() (CPUUsage, error) {
	g, err := cpu.Percent(1000*time.Millisecond, false)
	if err != nil {
		return CPUUsage{}, err
	}
	percent := g[0]

	e, err := cpu.Percent(1000*time.Millisecond, true)
	if err != nil {
		return CPUUsage{}, err
	}

	data := CPUUsage{
		GeneralPercent: percent,
		EachPercent:    e,
	}

	return data, nil
}

func getContainers() ([]docker.CgroupDockerStat, error) {
	d, err := docker.GetDockerStat()
	if err != nil {
		return nil, err
	}

	return d, nil
}
