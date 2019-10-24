package properties

import (
	"encoding/json"
	"net/url"
)

type ObjectProperty struct {
	ID   string   `json:"id"`
	Type string   `json:"type"`
	Name string   `json:"name"`
	IRI  *url.URL `json:"-"`
}

func (a *ObjectProperty) UnmarshalJSON(data []byte) error {
	var jsonData interface{}
	if err := json.Unmarshal(data, &jsonData); nil != err {
		return err
	}

	if iri, ok := jsonData.(string); ok {
		a.IRI, _ = url.Parse(iri)
		return nil
	}

	type alias ObjectProperty
	aa := &struct {
		*alias
	}{
		alias: (*alias)(a),
	}

	return json.Unmarshal(data, aa)
}

func (a *ObjectProperty) MarshalJSON() ([]byte, error) {
	if nil != a.IRI {
		return a.IRI.MarshalBinary()
	}

	return json.Marshal(a)
}
