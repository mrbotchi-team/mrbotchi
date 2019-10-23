package activitypub

import (
	"fmt"
	"net/http"

	"github.com/mrbotchi-team/mrbotchi/errors"
	"github.com/mrbotchi-team/mrbotchi/handler"
)

type (
	SharedInboxHandler struct {
		handler.HTTPHandler
	}
)

func (h SharedInboxHandler) Post(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Location", fmt.Sprintf("https://%s/accounts/%s/inbox", h.App.Config.Host, h.App.Config.User.Name))
	return errors.HTTPError{http.StatusPermanentRedirect, "Permanent Redirect"}
}
