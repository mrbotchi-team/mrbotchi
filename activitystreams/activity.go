package activitystreams

import "github.com/mrbotchi-team/mrbotchi/activitystreams/properties"

type Activity struct {
	Type   string                    `json:"type"`
	Actor  properties.ActorProperty  `json:"actor"`
	Object properties.ObjectProperty `json:"object"`
}
