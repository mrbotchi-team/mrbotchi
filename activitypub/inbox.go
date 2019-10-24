package activitypub

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mrbotchi-team/mrbotchi/activitystreams"
	"github.com/mrbotchi-team/mrbotchi/errors"
	"github.com/mrbotchi-team/mrbotchi/utils"

	"github.com/mrbotchi-team/mrbotchi/handler"
)

type (
	InboxHandler struct {
		handler.HTTPHandler
	}
)

func (h InboxHandler) Post(w http.ResponseWriter, r *http.Request) error {
	if !utils.RequestActivityJSON(r) {
		return errors.HTTPError{StatusCode: http.StatusBadRequest}
	}

	body, err := ioutil.ReadAll(r.Body)
	if nil != err {
		return errors.HTTPError{StatusCode: http.StatusBadRequest}
	}

	var activity activitystreams.Activity
	if err := json.Unmarshal(body, &activity); nil != err {
		return errors.HTTPError{StatusCode: http.StatusBadRequest}
	}

	switch activity.Type {
	case "Follow":
		return nil

	case "Undo":
		return nil

	default:
		return errors.HTTPError{StatusCode: http.StatusBadRequest}
	}
}
