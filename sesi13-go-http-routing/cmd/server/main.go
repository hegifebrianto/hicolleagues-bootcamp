package main

import (
	"log"
	"net/http"
	"time"

	"go-http-routing-sesi13/internal/handler"
	"go-http-routing-sesi13/internal/tugas"
)

func main() {
	repo := tugas.NewMemoryRepository()
	api := handler.New(repo)

	mux := http.NewServeMux()
	api.RegisterRoutes(mux)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("Server berjalan di http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}
