package app

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mrbotchi-team/mrbotchi/config"
	me "github.com/mrbotchi-team/mrbotchi/error"
	"github.com/mrbotchi-team/mrbotchi/handlers"
	"github.com/mrbotchi-team/mrbotchi/handlers/activitypub"
	"github.com/mrbotchi-team/mrbotchi/utils/response"
	"github.com/writeas/go-webfinger"
)

type (
	// App はTDNコンストラクタ。
	App struct {
		Router             *chi.Mux
		Config             *config.Config
		webFingerResolaver *webfingerResolaver
	}
	webfingerResolaver struct {
		host     string
		username string
	}
)

var (
	endpointNotfoundError = &me.APIError{ID: "ENDPOINT_NOTFOUND", Message: "This endpoint does not exist."}
	actorNotfoundError    = &me.APIError{ID: "ACTOR_NOTFOUND", Message: "This actor does not exist."}
)

// New は新しいAppを生成する関数。
func New(configfilePath string) (*App, error) {
	config, err := config.LoadConfig(configfilePath)
	if nil != err {
		return nil, err
	}

	return &App{
		Router: chi.NewRouter(),
		Config: config,
		webFingerResolaver: &webfingerResolaver{
			host:     config.Host,
			username: config.Actor.PreferredUsername,
		},
	}, nil
}

// Route はルーティングを行う関数。
func (a *App) Route() {
	a.Router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		response.WriteJSONResponse(w, http.StatusNotFound, endpointNotfoundError)
	})
	hs := handlerFactory()

	wr := webfinger.Default(a.webFingerResolaver)
	wr.NoTLSHandler = nil
	a.Router.Get(webfinger.WebFingerPath, http.HandlerFunc(wr.Webfinger))

	for endpoint, h := range hs {
		a.Router.Get(endpoint, handlers.HTTPHandlerFunc(h.Get).ServeHTTP)
		a.Router.Post(endpoint, handlers.HTTPHandlerFunc(h.Post).ServeHTTP)
		a.Router.Put(endpoint, handlers.HTTPHandlerFunc(h.Put).ServeHTTP)
		a.Router.Delete(endpoint, handlers.HTTPHandlerFunc(h.Delete).ServeHTTP)
	}
}

func handlerFactory() map[string]handlers.HTTPHandlerIF {
	var results map[string]handlers.HTTPHandlerIF = map[string]handlers.HTTPHandlerIF{
		// ActivityPub
		"/activitypub/inbox":     &activitypub.Inbox{HTTPHandler: handlers.HTTPHandler{}},
		"/activitypub/outbox":    &activitypub.Outbox{HTTPHandler: handlers.HTTPHandler{}},
		"/activitypub/followers": &activitypub.Followers{HTTPHandler: handlers.HTTPHandler{}},
		"/activitypub/following": &activitypub.Following{HTTPHandler: handlers.HTTPHandler{}},
		"/activitypub/liked":     &activitypub.Liked{HTTPHandler: handlers.HTTPHandler{}},
	}

	return results
}

func (wr webfingerResolaver) FindUser(username, hostname, requestHost string, r []webfinger.Rel) (*webfinger.Resource, error) {
	if hostname != wr.host || username != wr.username {
		return nil, actorNotfoundError
	}

	res := webfinger.Resource{
		Subject: "acct:" + username + "@" + hostname,
		Links: []webfinger.Link{
			{
				HRef: "https://" + hostname + "/",
				Type: "application/activity+json",
				Rel:  "Self",
			},
		},
	}
	return &res, nil
}

func (webfingerResolaver) DummyUser(username, hostname string, r []webfinger.Rel) (*webfinger.Resource, error) {
	return nil, actorNotfoundError
}

func (webfingerResolaver) IsNotFoundError(err error) bool {
	return err == actorNotfoundError
}
