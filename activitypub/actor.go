package activitypub

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/mrbotchi-team/mrbotchi/activitystreams/actor"
	"github.com/mrbotchi-team/mrbotchi/handler"
	"github.com/mrbotchi-team/mrbotchi/utils"
)

type (
	ActorHandler struct {
		handler.HTTPHandler
	}
)

func (h ActorHandler) Get(w http.ResponseWriter, r *http.Request) error {
	name := chi.URLParam(r, "name")
	if h.App.Config.User.Name != name {
		w.WriteHeader(http.StatusNotFound)
		return nil
	}

	person := actor.NewPerson(h.App.Config.Host, h.App.Config.User.Name, h.App.Config.User.DisplayName, "I'm Botchi.", h.App.Config.User.PublicKey)

	body, err := json.Marshal(person)
	if nil != err {
		return err
	}

	return utils.WriteBody(w, body, http.StatusOK, "application/activity+json; charset=utf8")
}
