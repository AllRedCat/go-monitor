package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"monitor/internal/services"
	"monitor/internal/utils"
)

func WSCpu(w http.ResponseWriter, r *http.Request) {
	conn, err := utils.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to set websocket upgrade:", err)
		return
	}
	defer conn.Close()

	for {
		cpuUsage, err := services.GetCpuUsage()
		if err != nil {
			log.Println("Failed to get CPU usage:", err)
			break
		}

		data, _ := json.Marshal(cpuUsage)
		if err := conn.WriteMessage(1, data); err != nil {
			log.Println("Failed to write message:", err)
			break
		}

		time.Sleep(3 * time.Second)
	}
}
