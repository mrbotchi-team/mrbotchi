package handlers

import (
	"fmt"
	"net/http"
)

type PingHandler struct {
	Handler
}

func (g PingHandler) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong!")
}
