package handlers

import (
	"net/http"
)

type (
	InboxHandler struct {
		Handler
	}
)

func (h InboxHandler) Post(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusNotImplemented)
	return nil
}
