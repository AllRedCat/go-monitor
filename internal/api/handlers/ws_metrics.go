package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"monitor/internal/services"
	"monitor/internal/utils"
)

func WSMetrics(w http.ResponseWriter, r *http.Request) {
	conn, err := utils.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to set websocket upgrade:", err)
		return
	}
	defer conn.Close()

	for {
		metrics, err := services.GetMetrics()
		if err != nil {
			log.Println("Failed to get metrics:", err)
			break
		}

		data, _ := json.Marshal(metrics)
		if err := conn.WriteMessage(1, data); err != nil {
			log.Println("Failed to write message:", err)
			break
		}

		time.Sleep(3 * time.Second)
	}
}