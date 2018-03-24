package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ryanyogan/k8-heart-beat/handlers"
	"github.com/ryanyogan/k8-heart-beat/version"
)

// PORT=8000[specify your own port number] go run main.go
func main() {
	log.Printf("Starting heart beat service...\ncommit: %s, build time: %s, release: %s",
		version.Commit, version.BuildTime, version.Release)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port was not set in environment variables.")
	}

	router := handlers.Router(version.BuildTime, version.Commit, version.Release)
	log.Print("The service is ready to listen and serve requests...")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
