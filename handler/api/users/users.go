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
		UserModel *models.UserModel
	}
	userRequest struct {
		Name     string `json:"name" validate:"required,printascii,min=1,max=16"`
		Password string `json:"password" validate:"required,printascii"`
	}
)

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

	h.UserModel.Insert(userBody.Name, userBody.Password, h.App.Config.Argon2)

	w.WriteHeader(http.StatusCreated)

	return nil
}
