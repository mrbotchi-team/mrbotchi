package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/mrbotchi-team/mrbotchi/activitystreams/actor"
	"github.com/mrbotchi-team/mrbotchi/error"
)

type (
	AccountHandler struct {
		Handler
	}
)

func (h AccountHandler) Get(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	if h.app.Config.User.Name != name {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	person := actor.NewPerson(h.app.Config.Host, h.app.Config.User.Name, h.app.Config.User.DisplayName, h.app.Config.User.PublicKey)

	w.Header().Set("content-type", "application/activity+json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(person); nil != err {
		error.NewInternalServerError().Response(w, r)
	}

}
