package activitystreams

import (
	"github.com/mrbotchi-team/mrbotchi/activitystreams/properties"
)

type Collection struct {
	ID      string                      `json:"id"`
	Type    string                      `json:"type"`
	Total   int                         `json:"totalItems"`
	Current properties.ObjectProperty   `json:"current,omitempty"`
	First   properties.ObjectProperty   `json:"first,omitempty"`
	Last    properties.ObjectProperty   `json:"last,omitempty"`
	Items   []properties.ObjectProperty `json:"items"`
}
