package handlers

import (
	"net/http"
)

type (
	OutboxHandler struct {
		Handler
	}
)

func (h OutboxHandler) Get(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusNotImplemented)
	return nil
}
