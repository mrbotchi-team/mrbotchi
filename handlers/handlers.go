package handlers

import (
	"net/http"

	"github.com/mrbotchi-team/mrbotchi/app"
	"github.com/mrbotchi-team/mrbotchi/error"
)

type (
	HandlerIf interface {
		Get(w http.ResponseWriter, r *http.Request)
		Post(w http.ResponseWriter, r *http.Request)
		Put(w http.ResponseWriter, r *http.Request)
		Delete(w http.ResponseWriter, r *http.Request)
	}
	Handler struct {
		app *app.App
	}
)

func (this Handler) Get(w http.ResponseWriter, r *http.Request) {
	error.NewMethodNotAllowed().Response(w, r)
}

func (this Handler) Post(w http.ResponseWriter, r *http.Request) {
	error.NewMethodNotAllowed().Response(w, r)
}

func (this Handler) Put(w http.ResponseWriter, r *http.Request) {
	error.NewMethodNotAllowed().Response(w, r)
}

func (this Handler) Delete(w http.ResponseWriter, r *http.Request) {
	error.NewMethodNotAllowed().Response(w, r)
}

func HandlerFactory(app *app.App) map[string]HandlerIf {
	var results map[string]HandlerIf = map[string]HandlerIf{
		"/{name}":        &AccountHandler{Handler{app}},
		"/{name}/inbox":  &InboxHandler{Handler{app}},
		"/{name}/outbox": &OutboxHandler{Handler{app}},
	}

	return results
}
