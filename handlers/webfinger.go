package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

type (
	WebfingerHandler struct {
		Handler
	}
	webfingerresponse struct {
		Subject string `json:"subject"`
		Links   []link `json:"links"`
	}
	link struct {
		Href       string             `json:"href"`
		Type       string             `json:"type,omitempty"`
		Rel        string             `json:"rel"`
		Properties map[string]*string `json:"properties,omitempty"`
		Titles     map[string]string  `json:"titles,omitempty"`
	}
)

func (h WebfingerHandler) Get(w http.ResponseWriter, r *http.Request) {
	uri := r.FormValue("resource")
	log.Println(uri)

	reg := regexp.MustCompile(`^acct:([a-zA-Z0-9_\-]+)@([a-zA-Z0-9_\-\.]+)`)
	subject := reg.FindSubmatch([]byte(uri))
	if 0 >= len(subject) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	account := string(subject[1])
	host := string(subject[2])

	if h.app.Config.User.Name != account || h.app.Config.Host != host {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	href := fmt.Sprintf("https://%s/%s", h.app.Config.Host, account)

	response := webfingerresponse{
		Subject: string(subject[0]),
		Links: []link{
			link{
				Rel:  "self",
				Type: "application/activity+json",
				Href: href,
			},
		},
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); nil != err {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
