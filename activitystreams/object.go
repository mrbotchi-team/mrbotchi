package activitystreams

type Object struct {
	Context []string `json:"@context"`
	Type    string   `json:"type"`
	ID      string   `json:"id"`
	Name    string   `json:"name"`
}

func NewObject(context []string, object_type, id, name string) *Object {
	context = append(context, "https://www.w3.org/ns/activitystreams")

	return &Object{
		Context: context,
		Type:    object_type,
		ID:      id,
		Name:    name,
	}
}
