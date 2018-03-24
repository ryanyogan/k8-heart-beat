package handlers

import (
	"fmt"
	"net/http"
)

func ping(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Pong!")
}
