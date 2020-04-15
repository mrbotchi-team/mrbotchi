package app

import (
	"github.com/go-chi/chi"
	"github.com/mrbotchi-team/mrbotchi/handlers"
	"github.com/mrbotchi-team/mrbotchi/handlers/activitypub"
)

// App はTDNコンストラクタ。
type App struct {
	Router *chi.Mux
}

// New は新しいAppを生成する関数。
func New() *App {
	return &App{
		Router: chi.NewRouter(),
	}
}

// Route はルーティングを行う関数。
func (a *App) Route() {
	hs := handlerFactory()

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
