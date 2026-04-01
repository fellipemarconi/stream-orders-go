// Package main is the main package for the API server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"stream-orders/internal/handler"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	orderHandler := handler.NewOrderHandler()

	mux := http.NewServeMux()
	
	mux.HandleFunc("/health", orderHandler.HealthCheck)
	
	mux.HandleFunc("/api/orders", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			orderHandler.ListOrders(w, r)
		case http.MethodPost:
			orderHandler.CreateOrder(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	addr := fmt.Sprintf(":%s", port)
	log.Printf("🚀 API Server starting on %s", addr)
	
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
