package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/mrbotchi-team/mrbotchi/errors"
)

func WriteBody(w http.ResponseWriter, body []byte, status int, contentType string) error {
	w.Header().Set("Content-Type", contentType+"; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(body)))
	w.WriteHeader(status)
	_, err := w.Write(body)

	return err
}

func WriteError(w http.ResponseWriter, e error) {
	if err, ok := e.(errors.APIError); ok {
		body, _ := json.Marshal(err)
		WriteBody(w, body, err.StatusCode, "application/json")
		return
	}
	if err, ok := e.(errors.HTTPError); ok {
		WriteBody(w, []byte(err.Message), err.StatusCode, "text/plain")
		return
	}

	log.Println("Error:", e)

	err := errors.InternalServerError()
	body, _ := json.Marshal(err)
	WriteBody(w, body, err.StatusCode, "application/json")
}
