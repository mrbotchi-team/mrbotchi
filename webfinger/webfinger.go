package webfinger

import (
	"fmt"
	"net/http"

	"github.com/mrbotchi-team/mrbotchi/errors"
	"github.com/writeas/go-webfinger"
)

type WebfingerResolver struct {
	UserName string
	Host     string
}

func (resolver WebfingerResolver) FindUser(username, hostname, requestHost string, r []webfinger.Rel) (*webfinger.Resource, error) {
	if resolver.UserName != username || resolver.Host != hostname {
		return nil, errors.HTTPError{StatusCode: http.StatusNotFound}
	}

	res := webfinger.Resource{
		Subject: "acct:" + username + "@" + hostname,
		Links: []webfinger.Link{
			{
				HRef: fmt.Sprintf("https://%s/accounts/%s", hostname, username),
				Type: "application/activity+json",
				Rel:  "self",
			},
		},
	}

	return &res, nil
}

func (resolver WebfingerResolver) DummyUser(username string, hostname string, r []webfinger.Rel) (*webfinger.Resource, error) {
	return nil, errors.HTTPError{StatusCode: http.StatusNotFound}
}

func (resolver WebfingerResolver) IsNotFoundError(err error) bool {
	return err == errors.HTTPError{StatusCode: http.StatusNotFound}
}
