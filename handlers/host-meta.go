package handlers

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/mrbotchi-team/mrbotchi/error"
)

type (
	hostMeta struct {
		XMLName xml.Name `xml:"XRD"`
		Text    string   `xml:",chardata"`
		Xmlns   string   `xml:"xmlns,attr"`
		Link    struct {
			Text     string `xml:",chardata"`
			Rel      string `xml:"rel,attr"`
			Type     string `xml:"type,attr"`
			Template string `xml:"template,attr"`
		} `xml:"Link"`
	}
	HostMetaHandler struct {
		Handler
	}
)

func newHostmeta(host string) *hostMeta {
	return &hostMeta{
		Xmlns: "http://docs.oasis-open.org/ns/xri/xrd-1.0",
		Link: struct {
			Text     string `xml:",chardata"`
			Rel      string `xml:"rel,attr"`
			Type     string `xml:"type,attr"`
			Template string `xml:"template,attr"`
		}{
			Rel:      "lrdd",
			Type:     "application/xrd+xml",
			Template: fmt.Sprintf("https://%s/.well-known/webfinger?resource={uri}", host),
		},
	}
}

func (h HostMetaHandler) Get(w http.ResponseWriter, r *http.Request) {
	response := newHostmeta(fmt.Sprintf("%s:%v", h.app.Config.Host, h.app.Config.Port))

	w.Header().Set("content-type", "application/xrd+xml")
	w.WriteHeader(http.StatusOK)

	if err := xml.NewEncoder(w).Encode(response); nil != err {
		error.NewInternalServerError().Response(w, r)
	}
}
