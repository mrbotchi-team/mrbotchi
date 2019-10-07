package handler

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
	HTTPHandler struct {
		App *app.App
	}
	HandlerFunc func(http.ResponseWriter, *http.Request) error
)

func (h HTTPHandler) Get(w http.ResponseWriter, r *http.Request) error {
	return errors.MethodNotAllowed()
}

func (h HTTPHandler) Post(w http.ResponseWriter, r *http.Request) error {
	return errors.MethodNotAllowed()
}

func (h HTTPHandler) Put(w http.ResponseWriter, r *http.Request) error {
	return errors.MethodNotAllowed()
}

func (h HTTPHandler) Delete(w http.ResponseWriter, r *http.Request) error {
	return errors.MethodNotAllowed()
}

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h(w, r)
	if nil != err {
		utils.WriteError(w, errors.InternalServerError())
	}
}
