package handlers

import (
	"log"
	"sync/atomic"
	"time"

	"github.com/gorilla/mux"
)

// Router creates the root router for the app which writes to r
func Router(buildTime, commit, release string) *mux.Router {
	isReady := &atomic.Value{}
	isReady.Store(false)

	go func() {
		log.Printf("Readyz probe is negative by default...")
		// I like to let the cache warm up for a few seconds
		time.Sleep(5 * time.Second)
		isReady.Store(true)
		log.Printf("Readyz probe is positive.")
	}()

	r := mux.NewRouter()
	r.HandleFunc("/ping", ping(buildTime, commit, release)).Methods("GET")
	r.HandleFunc("/healthz", healthz)
	r.HandleFunc("/readyz", readyz(isReady))
	return r
}
