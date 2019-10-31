package activitystreams

import (
	"github.com/mrbotchi-team/mrbotchi/activitystreams/properties"
)

type Collection struct {
	ID      string                      `json:"id"`
	Type    string                      `json:"type"`
	Total   int                         `json:"totalItems"`
	Current properties.ObjectProperty   `json:"current,omitempt"`
	First   properties.ObjectProperty   `json:"first,omitempt"`
	Last    properties.ObjectProperty   `json:"last,omitempt"`
	Items   []properties.ObjectProperty `json:"items"`
}
