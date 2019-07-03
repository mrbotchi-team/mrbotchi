package handlers

import (
	"net/http"
)

type (
	HandlerIf interface {
		Get(w http.ResponseWriter, r *http.Request)
		Post(w http.ResponseWriter, r *http.Request)
		Put(w http.ResponseWriter, r *http.Request)
		Delete(w http.ResponseWriter, r *http.Request)
	}
	Handler struct {
	}
)

func (this Handler) Get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (this Handler) Post(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (this Handler) Put(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (this Handler) Delete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func HandlerFactory() map[string]HandlerIf {
	var results map[string]HandlerIf = map[string]HandlerIf{
		"/ping": &PingHandler{},
	}

	return results
}
