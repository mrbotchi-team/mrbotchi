package handler

import (
	"net/http"

	"github.com/mrbotchi-team/mrbotchi/errors"
)

type (
	SchwimmwagenHandler struct {
		HTTPHandler
	}
)

func (h SchwimmwagenHandler) Get(w http.ResponseWriter, r *http.Request) error {
	return errors.ImSchwimmwagen()
}
