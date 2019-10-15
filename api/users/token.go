package users

import (
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

	var user models.User
	h.App.DB.Get(&user, "SELECT * FROM users WHERE name=$1", userName)

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
