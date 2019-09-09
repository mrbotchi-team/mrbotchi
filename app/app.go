package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mrbotchi-team/mrbotchi/webfinger"
	wf "github.com/writeas/go-webfinger"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mrbotchi-team/mrbotchi/config"
	"github.com/mrbotchi-team/mrbotchi/error"
)

type App struct {
	Router  *chi.Mux
	Config  *config.Config
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

	config := config.LoadConfig()
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable", config.DB.Host, config.DB.Port, config.DB.User, config.DB.Password, config.DB.DBname))
	if nil != err {
		log.Fatalln(err)
	}

	webfinger := wf.Default(webfinger.WebfingerResolver{UserName: config.User.Name, Host: config.Host})
	webfinger.NoTLSHandler = nil

	router.Get(wf.WebFingerPath, http.HandlerFunc(webfinger.Webfinger))

	return &App{
		Router:  router,
		Config:  config,
		DB:      db,
		Version: version,
	}
}
