package users

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/mrbotchi-team/mrbotchi/utils"

	"github.com/go-chi/chi"
	"github.com/o1egl/paseto"

	"github.com/mrbotchi-team/mrbotchi/handler"
	"github.com/mrbotchi-team/mrbotchi/models"

	"github.com/mrbotchi-team/mrbotchi/errors"
)

type (
	TokenHandler struct {
		handler.HTTPHandler
		UserModel *models.UserModel
	}
)

func (h TokenHandler) Get(w http.ResponseWriter, r *http.Request) error {
	v := r.URL.Query()
	if nil == v {
		return errors.InvalidRequest()
	}

	password := v.Get("password")
	if "" == password {
		return errors.InvalidRequest()
	}

	userName := chi.URLParam(r, "name")

	user := h.UserModel.FindByName(userName)

	ok, err := utils.VerifyPassword(password, user.Password)
	if nil != err {
		return err
	}
	if ok {
		symmetricKey := []byte(h.App.Config.PasetoKey)
		now := time.Now()
		exp := now.Add(24 * time.Hour)
		nbt := now

		token := paseto.JSONToken{
			Subject:    userName,
			IssuedAt:   now,
			Expiration: exp,
			NotBefore:  nbt,
		}

		response, err := paseto.NewV2().Encrypt(symmetricKey, token, nil)
		if nil != err {
			return err
		}
		utils.WriteBody(w, []byte(response), http.StatusOK, "text/plain")
	} else {
		return errors.AuthFailed()
	}

	return nil
}

func (h TokenHandler) Put(w http.ResponseWriter, r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if nil != err {
		return errors.InvalidRequest()
	}
	defer r.Body.Close()

	userName := chi.URLParam(r, "name")

	symmetricKey := []byte(h.App.Config.PasetoKey)

	var token paseto.JSONToken
	if err := paseto.NewV2().Decrypt(string(body), symmetricKey, &token, nil); nil != err {
		log.Println(err)
		return errors.InvalidRequest()
	}

	if time.Now().After(token.Expiration) {
		return errors.InvalidRequest()
	}

	if userName != token.Subject {
		return errors.InvalidRequest()
	}

	now := time.Now()
	exp := now.Add(24 * time.Hour)
	nbt := now

	newToken := paseto.JSONToken{
		Subject:    userName,
		IssuedAt:   now,
		Expiration: exp,
		NotBefore:  nbt,
	}

	response, err := paseto.NewV2().Encrypt(symmetricKey, newToken, nil)
	if nil != err {
		return err
	}
	utils.WriteBody(w, []byte(response), http.StatusOK, "text/plain")

	return nil
}
