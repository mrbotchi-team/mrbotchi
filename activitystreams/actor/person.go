package actor

import (
	"fmt"
	"strings"

	"github.com/mrbotchi-team/mrbotchi/activitystreams"
)

type (
	Person struct {
		*activitystreams.Object
		Following                 string   `json:"following"`
		Followers                 string   `json:"followers"`
		Liked                     string   `json:"liked"`
		Inbox                     string   `json:"inbox"`
		Outbox                    string   `json:"outbox"`
		Endpoints                 Endpoint `json:"endpoints"`
		PreferredUsername         string   `json:"preferredUsername"`
		Name                      string   `json:"name"`
		Summary                   string   `json:"summary"`
		ManuallyApprovesFollowers bool     `json:"manuallyApprovesFollowers"`
		PublicKey                 PubKey   `json:"publicKey"`
	}
	Endpoint struct {
		SharedInbox string `json:"sharedInbox"`
	}
	PubKey struct {
		ID           string `json:"id"`
		Type         string `json:"type"`
		Owner        string `json:"owner"`
		PublicKeyPem string `json:"publicKeyPem"`
	}
)

func NewPerson(host, name, display_name, summary, publicKey string) *Person {
	id := fmt.Sprintf("https://%s/%s", host, name)
	following := strings.Join([]string{id, "/following"}, "")
	followers := strings.Join([]string{id, "/followers"}, "")
	liked := strings.Join([]string{id, "/liked"}, "")
	inbox := strings.Join([]string{id, "/inbox"}, "")
	outbox := strings.Join([]string{id, "/outbox"}, "")
	publicKeyID := strings.Join([]string{id, "/publickey"}, "")

	object := activitystreams.NewObject([]string{"https://w3id.org/security/v1"}, "Person", id, name)

	return &Person{
		Object:                    object,
		Following:                 following,
		Followers:                 followers,
		Liked:                     liked,
		Inbox:                     inbox,
		Outbox:                    outbox,
		Endpoints:                 Endpoint{SharedInbox: inbox},
		PreferredUsername:         name,
		ManuallyApprovesFollowers: false,
		Name:                      display_name,
		Summary:                   summary,
		PublicKey:                 PubKey{ID: publicKeyID, Type: "Key", Owner: id, PublicKeyPem: publicKey},
	}
}
