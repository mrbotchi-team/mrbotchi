package users

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mrbotchi-team/mrbotchi/errors"
	"github.com/mrbotchi-team/mrbotchi/models"
	"github.com/mrbotchi-team/mrbotchi/utils"

	"github.com/mrbotchi-team/mrbotchi/handler"
)

type (
	UsersHandler struct {
		handler.HTTPHandler
	}
	userRequest struct {
		Name     string `json:"name" validate:"required,printascii,min=1,max=16"`
		Password string `json:"password" validate:"required,printascii"`
	}
)

func (h UsersHandler) Get(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h UsersHandler) Post(w http.ResponseWriter, r *http.Request) error {
	if !utils.RequestJSON(r) {
		return errors.InvalidCT("application/json")
	}

	body, err := ioutil.ReadAll(r.Body)
	if nil != err {
		return errors.InvalidRequest()
	}
	defer r.Body.Close()

	var userBody userRequest
	if err := json.Unmarshal(body, &userBody); nil != err {
		return errors.InvalidRequest()
	}

	if err := h.App.Validate.Struct(userBody); nil != err {
		return errors.InvalidRequest()
	}

	user, err := models.NewUser(userBody.Name, userBody.Password, h.App.Config.Argon2)
	if nil != err {
		return err
	}

	h.App.DB.MustExec("INSERT INTO users (name, password, is_deleted, created_at) VALUES ($1, $2, $3, $4)", user.Name, user.Password, user.IsDeleted, user.CreatedAt)

	w.WriteHeader(http.StatusCreated)

	return nil
}
