package handler

import (
	"log"
	"net/http"

	"github.com/mrbotchi-team/mrbotchi/utils"
	"gopkg.in/spacemonkeygo/httpsig.v0"

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
		App    *app.App
		Signer *httpsig.Signer
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
		log.Println("Oops... Something worng...")
		log.Println("Error:", err)

		if err, ok := err.(errors.APIError); ok {
			utils.WriteError(w, err)
			return
		}

		utils.WriteError(w, errors.InternalServerError())
	}
}

func NewHandler(app *app.App, signer *httpsig.Signer) HTTPHandler {
	return HTTPHandler{
		App:    app,
		Signer: signer,
	}
}
