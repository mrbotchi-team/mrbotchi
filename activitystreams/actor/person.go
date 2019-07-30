package actor

import "fmt"

type Person struct {
	Context           []string `json:"@context"`
	ID                string   `json:"id"`
	Type              string   `json:"type"`
	PreferredUserName string   `json:"preferredUsername"`
	Inbox             string   `json:"inbox"`
	Outbox            string   `json:"outbox"`
	PublicKey         struct {
		ID           string `json:"id"`
		Owner        string `json:"owner"`
		PublicKeyPem string `json:"publicKeyPem"`
	} `json:"publickey"`
}

func NewPerson(host, name, preferred_username, publickey string) *Person {
	id := fmt.Sprintf("https://%s/%s", host, name)

	return &Person{
		Context: []string{
			"https://www.w3.org/ns/activitystreams",
			"https://w3id.org/security/v1",
		},
		ID:                id,
		Type:              "Person",
		PreferredUserName: preferred_username,
		Inbox:             fmt.Sprintf("%s/inbox", id),
		Outbox:            fmt.Sprintf("%s/outbox", id),
		PublicKey: struct {
			ID           string `json:"id"`
			Owner        string `json:"owner"`
			PublicKeyPem string `json:"publicKeyPem"`
		}{
			ID:           fmt.Sprintf("%s/pubkey", id),
			Owner:        id,
			PublicKeyPem: publickey,
		},
	}
}
