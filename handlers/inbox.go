package handlers

import (
	"net/http"
)

type (
	InboxHandler struct {
		Handler
	}
)

func (h InboxHandler) Post(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
