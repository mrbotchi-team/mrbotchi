package activitypub

import (
	"net/http"

	"github.com/mrbotchi-team/mrbotchi/handler"
)

type (
	InboxHandler struct {
		handler.HTTPHandler
	}
)

func (h InboxHandler) Post(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusNotImplemented)
	return nil
}
