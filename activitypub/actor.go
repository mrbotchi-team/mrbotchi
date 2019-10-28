package activitypub

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"

	"github.com/mrbotchi-team/mrbotchi/activitystreams"
	"github.com/mrbotchi-team/mrbotchi/handler"
	"github.com/mrbotchi-team/mrbotchi/securityvocabulary"
	"github.com/mrbotchi-team/mrbotchi/utils"
)

type (
	ActorHandler struct {
		handler.HTTPHandler
	}
	actorResponse struct {
		*activitystreams.Actor
		Following                 string                        `json:"following"`
		Followers                 string                        `json:"followers"`
		Liked                     string                        `json:"liked"`
		Inbox                     string                        `json:"inbox"`
		Outbox                    string                        `json:"outbox"`
		PreferredUsername         string                        `json:"preferredUsername"`
		Summary                   string                        `json:"summary"`
		ManuallyApprovesFollowers bool                          `json:"manuallyApprovesFollowers"`
		PublicKey                 *securityvocabulary.PublicKey `json:"publicKey"`
	}
)

func (h ActorHandler) Get(w http.ResponseWriter, r *http.Request) error {
	name := chi.URLParam(r, "name")
	if h.App.Config.Account.Name != name {
		w.WriteHeader(http.StatusNotFound)
		return nil
	}

	id := fmt.Sprintf("https://%s/%s", h.App.Config.Host, name)
	following := strings.Join([]string{id, "following"}, "/")
	followers := strings.Join([]string{id, "followers"}, "/")
	liked := strings.Join([]string{id, "liked"}, "/")
	inbox := strings.Join([]string{id, "inbox"}, "/")
	outbox := strings.Join([]string{id, "outbox"}, "/")

	publicKeyEndpoint := strings.Join([]string{id, "publickey"}, "/")
	publickey := securityvocabulary.NewPublicKey(publicKeyEndpoint, id, h.App.Config.Account.PublicKey)

	person := actorResponse{
		Actor: &activitystreams.Actor{
			ID:   id,
			Type: "Person",
			Name: name,
		},
		Following:                 following,
		Followers:                 followers,
		Liked:                     liked,
		Inbox:                     inbox,
		Outbox:                    outbox,
		PreferredUsername:         h.App.Config.Account.DisplayName,
		Summary:                   h.App.Config.Account.Summary,
		ManuallyApprovesFollowers: false,
		PublicKey:                 publickey,
	}

	type alias actorResponse
	p := &struct {
		Context []string `json:"@context"`
		*alias
	}{
		Context: []string{
			"https://www.w3.org/ns/activitystreams",
			"https://w3id.org/security/v1",
		},
		alias: (*alias)(&person),
	}

	body, err := json.Marshal(p)
	if nil != err {
		return err
	}

	return utils.WriteBody(w, body, http.StatusOK, "application/activity+json; charset=utf8")
}
