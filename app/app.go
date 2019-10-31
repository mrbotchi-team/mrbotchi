package app

import (
	"net/http"

	"github.com/mrbotchi-team/mrbotchi/errors"
	"github.com/mrbotchi-team/mrbotchi/utils"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mrbotchi-team/mrbotchi/config"
	validate "gopkg.in/go-playground/validator.v9"
)

type App struct {
	Router   *chi.Mux
	Config   *config.Config
	Validate *validate.Validate
}

func NewApp() *App {
	router := chi.NewRouter()
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.GetHead)

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		utils.WriteError(w, errors.EndpointNotFoundError())
	})

	config := config.LoadConfig()

	validate := validate.New()

	return &App{
		Router:   router,
		Config:   config,
		Validate: validate,
	}
}
