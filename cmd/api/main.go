package main

import (
	"log"
	"net/http"

	"monitor/internal/api"
)

func main() {
	router := api.NewRouter()

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("ListenAndServe error:", err)
	}
}
