package securityvocabulary

type PublicKey struct {
	Type         string `json:"type"`
	ID           string `json:"id"`
	Owner        string `json:"owner"`
	PublicKeyPem string `json:"publicKeyPem"`
}

func NewPublicKey(endpoint, owner, publickey string) *PublicKey {

	return &PublicKey{
		Type:         "Key",
		ID:           endpoint,
		Owner:        owner,
		PublicKeyPem: publickey,
	}
}
