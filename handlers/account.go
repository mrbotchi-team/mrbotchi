package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/mrbotchi-team/mrbotchi/activitystreams/actor"
	"github.com/mrbotchi-team/mrbotchi/utils"
)

type (
	AccountHandler struct {
		Handler
	}
)

func (h AccountHandler) Get(w http.ResponseWriter, r *http.Request) error {
	name := chi.URLParam(r, "name")
	if h.app.Config.User.Name != name {
		w.WriteHeader(http.StatusNotFound)
		return nil
	}

	person := actor.NewPerson(h.app.Config.Host, h.app.Config.User.Name, h.app.Config.User.DisplayName, h.app.Config.User.PublicKey)

	body, err := json.Marshal(person)
	if nil != err {
		return err
	}

	return utils.WriteBody(w, body, http.StatusOK, "application/activity+json")
}
