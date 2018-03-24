package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Pong!")
	})
	http.ListenAndServe(":8000", nil)
}
