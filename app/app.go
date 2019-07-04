package app

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mr-botchi/backend/error"
)

type App struct {
	Router  *chi.Mux
	Version string
}

func NewApp(version string) *App {
	router := chi.NewRouter()
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		error.NewEndpointNotFoundError().Response(w, r)
	})

	return &App{
		Router:  router,
		Version: version,
	}
}
