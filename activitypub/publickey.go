package activitypub

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"

	"github.com/mrbotchi-team/mrbotchi/handler"
	"github.com/mrbotchi-team/mrbotchi/securityvocabulary"
	"github.com/mrbotchi-team/mrbotchi/utils"
)

type (
	PublickeyHandler struct {
		handler.HTTPHandler
	}
)

func (h PublickeyHandler) Get(w http.ResponseWriter, r *http.Request) error {
	name := chi.URLParam(r, "name")
	if h.App.Config.User.Name != name {
		w.WriteHeader(http.StatusNotFound)
		return nil
	}

	id := fmt.Sprintf("https://%s/accounts/%s", h.App.Config.Host, name)
	endpoint := strings.Join([]string{id, "/publickey"}, "")
	publickey := securityvocabulary.NewPublicKey(endpoint, id, h.App.Config.User.PublicKey)

	type alias securityvocabulary.PublicKey
	p := &struct {
		Context string `json:"@context"`
		*alias
	}{
		Context: "https://w3id.org/security/v1",
		alias:   (*alias)(publickey),
	}

	body, err := json.Marshal(p)
	if nil != err {
		return err
	}

	return utils.WriteBody(w, body, http.StatusOK, "application/ld+json; charset=utf8")
}
