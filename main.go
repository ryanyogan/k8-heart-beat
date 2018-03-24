package main

import (
	"log"
	"net/http"

	"github.com/ryanyogan/k8-heart-beat/handlers"
)

func main() {
	log.Print("Starting heart beat service...")
	router := handlers.Router()
	log.Print("The service is ready to listen and serve requests...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
