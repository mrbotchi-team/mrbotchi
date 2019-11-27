package activitypub

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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
		PreferredUsername         string                        `json:"preferredUsername"`
		Summary                   string                        `json:"summary"`
		ManuallyApprovesFollowers bool                          `json:"manuallyApprovesFollowers"`
		PublicKey                 *securityvocabulary.PublicKey `json:"publicKey"`
	}
)

func (h ActorHandler) Get(w http.ResponseWriter, r *http.Request) error {
	id := fmt.Sprintf("https://%s", h.App.Config.Host)
	following := strings.Join([]string{id, "following"}, "/")
	followers := strings.Join([]string{id, "followers"}, "/")
	liked := strings.Join([]string{id, "liked"}, "/")
	inbox := strings.Join([]string{id, "inbox"}, "/")
	outbox := strings.Join([]string{id, "outbox"}, "/")

	publicKeyEndpoint := strings.Join([]string{id, "publickey"}, "/")
	publickey := securityvocabulary.NewPublicKey(publicKeyEndpoint, id, h.App.Config.Account.PublicKey)

	person := actorResponse{
		Actor: &activitystreams.Actor{
			ID:     id,
			Type:   "Person",
			Name:   h.App.Config.Account.DisplayName,
			Inbox:  inbox,
			Outbox: outbox,
		},
		Following:                 following,
		Followers:                 followers,
		Liked:                     liked,
		PreferredUsername:         h.App.Config.Account.Name,
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
