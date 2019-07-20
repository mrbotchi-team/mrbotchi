package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mr-botchi/backend/error"
)

type App struct {
	Router  *chi.Mux
	Config  *Config
	DB      *sqlx.DB
	Version string
}

func NewApp(version string) *App {
	router := chi.NewRouter()
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.GetHead)

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		error.NewEndpointNotFoundError().Response(w, r)
	})

	config := loadConfig()
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable", config.DB.Host, config.DB.Port, config.DB.User, config.DB.Password, config.DB.DBname))
	if nil != err {
		log.Fatalln(err)
	}

	return &App{
		Router:  router,
		Config:  config,
		DB:      db,
		Version: version,
	}
}
