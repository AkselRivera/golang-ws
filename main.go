package main

import (
	"log"
	"net/http"
	"ws/cmd/web"
	"ws/internal/handlers"
)

func main() {
	mux := web.Routes()

	log.Println("Starting channel listener...")
	go handlers.ListenToWsChannel()

	log.Println("Staring web server on port 8080")

	_ = http.ListenAndServe(":8080", mux)
}
