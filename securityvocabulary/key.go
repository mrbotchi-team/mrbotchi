package securityvocabulary

type Key struct {
	Context      []string `json:"@context"`
	Type         string   `json:"type"`
	ID           string   `json:"id"`
	Owner        string   `json:"owner"`
	PublicKeyPem string   `json:"publicKeyPem"`
}

func NewKey(endpoint, owner, publickey string) *Key {
	context := []string{"https://w3id.org/security/v1"}

	return &Key{
		Context:      context,
		Type:         "Key",
		ID:           endpoint,
		Owner:        owner,
		PublicKeyPem: publickey,
	}
}
