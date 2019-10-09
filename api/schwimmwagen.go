package api

import (
	"net/http"

	handler "github.com/mrbotchi-team/mrbotchi/handler"

	"github.com/mrbotchi-team/mrbotchi/errors"
)

type (
	SchwimmwagenHandler struct {
		handler.HTTPHandler
	}
)

func (h SchwimmwagenHandler) Get(w http.ResponseWriter, r *http.Request) error {
	return errors.ImSchwimmwagen()
}
