package utils

import (
	"mime"
	"net/http"
)

func RequestJSON(r *http.Request) bool {
	ct, _, _ := mime.ParseMediaType(r.Header.Get("Content-Type"))
	return "application/json" == ct
}

func RequestActivityJSON(r *http.Request) bool {
	ct, params, _ := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if "application/activity+json" == ct {
		return true
	} else if "application/ld+json" == ct {
		if param, ok := params["profile"]; ok {
			return "https://www.w3.org/ns/activitystreams" == param
		}
	}

	return false
}
