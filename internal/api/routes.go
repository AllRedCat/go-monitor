// Package api
package api

import (
	"net/http"

	"monitor/internal/api/handlers"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// REST Routes
	mux.HandleFunc("/info", handlers.GetSystemInfo)
	mux.HandleFunc("/docker", handlers.Docker)
	mux.HandleFunc("/exec", handlers.ExecHandler)

	// WebSocket Routes
	mux.HandleFunc("/ws/info", handlers.WSMetrics)
	mux.HandleFunc("/ws/cpu", handlers.WSCpu)

	return mux
}
