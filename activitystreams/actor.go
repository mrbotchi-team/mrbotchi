package activitystreams

import "github.com/mrbotchi-team/mrbotchi/activitystreams/properties"

type Actor struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Name   string `json:"name"`
	Inbox  string `json:"inbox"`
	Outbox string `json:"outbox"`
}

func (a *Actor) ToActorProperties() *properties.ActorProperty {
	return &properties.ActorProperty{
		ID:   a.ID,
		Type: a.Type,
		Name: a.Name,
	}
}
