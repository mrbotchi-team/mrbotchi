package utils

import (
	"mime"
	"net/http"
)

func RequestJSON(r *http.Request) bool {
	ct, _, _ := mime.ParseMediaType(r.Header.Get("Content-Type"))
	return "application/json" == ct
}
