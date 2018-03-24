package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	log.Print("The service is ready to listen and serve requests...")

	killSignal := <-interrupt
	switch killSignal {
	case os.Kill:
		log.Print("Got SIGKILL...")
	case os.Interrupt:
		log.Print("Got SIGINT...")
	case syscall.SIGTERM:
		log.Print("Got SIGTERM...")
	}

	log.Print("The service is shutting down...")
	srv.Shutdown(context.Background())
	log.Print("Done")

}
