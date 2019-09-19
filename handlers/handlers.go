package handlers

import (
	"net/http"

	"github.com/mrbotchi-team/mrbotchi/utils"

	"github.com/mrbotchi-team/mrbotchi/app"
	"github.com/mrbotchi-team/mrbotchi/errors"
)

type (
	HandlerIf interface {
		Get(w http.ResponseWriter, r *http.Request) error
		Post(w http.ResponseWriter, r *http.Request) error
		Put(w http.ResponseWriter, r *http.Request) error
		Delete(w http.ResponseWriter, r *http.Request) error
	}
	Handler struct {
		app *app.App
	}
	HandlerFunc func(http.ResponseWriter, *http.Request) error
)

func (h Handler) Get(w http.ResponseWriter, r *http.Request) error {
	return errors.MethodNotAllowed()
}

func (h Handler) Post(w http.ResponseWriter, r *http.Request) error {
	return errors.MethodNotAllowed()
}

func (h Handler) Put(w http.ResponseWriter, r *http.Request) error {
	return errors.MethodNotAllowed()
}

func (h Handler) Delete(w http.ResponseWriter, r *http.Request) error {
	return errors.MethodNotAllowed()
}

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h(w, r)
	if nil != err {
		utils.WriteError(w, errors.InternalServerError())
	}
}

func HandlerFactory(app *app.App) map[string]HandlerIf {
	var results map[string]HandlerIf = map[string]HandlerIf{
		"/{name}":        &AccountHandler{Handler{app}},
		"/{name}/inbox":  &InboxHandler{Handler{app}},
		"/{name}/outbox": &OutboxHandler{Handler{app}},
	}

	return results
}
