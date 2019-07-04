package error

import (
	"encoding/json"
	"net/http"
)

type (
	ErrorIf interface {
		Response(w http.ResponseWriter, r *http.Request)
	}
	Error struct {
		StatusCode int    `json:"status_code"`
		ErrorID    string `json:"id"`
		Message    string `json:"message"`
	}
)

func (e Error) Response(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.StatusCode)

	//nolint
	json.NewEncoder(w).Encode(e)
}
