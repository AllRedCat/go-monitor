// Package handlers
package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"monitor/internal/services"
)

func Docker(w http.ResponseWriter, r *http.Request) {
	// Check if the Method it's GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	doc, err := services.GetContainers() // Cal function to get containers info
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
}
