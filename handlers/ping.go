package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mr-botchi/backend/error"
)

type (
	PingHandler struct {
		Handler
	}
	pingResponse struct {
		Message string `json:"message"`
		Version string `json:"version"`
	}
)

func (h PingHandler) Get(w http.ResponseWriter, r *http.Request) {
	response := pingResponse{Message: "Guten Morgen!! I'm Mr Botchi! maybe backend!", Version: h.app.Version}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); nil != err {
		error.NewInternalServerError().Response(w, r)
	}
}
