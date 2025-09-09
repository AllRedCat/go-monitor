package handlers

import (
	"encoding/json"
	"net/http"

	"monitor/internal/models"
	"monitor/internal/services"
)

func ExecHandler(w http.ResponseWriter, r *http.Request) {
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
	var req models.ExecRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Path == "" {
		http.Error(w, "\"path\" is required", http.StatusBadRequest)
		return
	}

	resp, err := services.RunScript(req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
