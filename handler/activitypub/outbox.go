package activitypub

import (
	"net/http"

	"github.com/mrbotchi-team/mrbotchi/handler"
)

type (
	OutboxHandler struct {
		handler.HTTPHandler
	}
)

func (h OutboxHandler) Get(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusNotImplemented)
	return nil
}
