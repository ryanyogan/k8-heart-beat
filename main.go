package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ryanyogan/k8-heart-beat/handlers"
)

// PORT=8000[specify your own port number] go run main.go
func main() {
	log.Print("Starting heart beat service...")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port was not set in environment variables.")
	}

	router := handlers.Router()
	log.Print("The service is ready to listen and serve requests...")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
