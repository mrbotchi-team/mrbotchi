package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	wf "github.com/writeas/go-webfinger"

	"github.com/mrbotchi-team/mrbotchi/app"
	"github.com/mrbotchi-team/mrbotchi/handler"
	"github.com/mrbotchi-team/mrbotchi/handler/activitypub"
	"github.com/mrbotchi-team/mrbotchi/handler/api"
	"github.com/mrbotchi-team/mrbotchi/handler/api/users"
	"github.com/mrbotchi-team/mrbotchi/models"
	"github.com/mrbotchi-team/mrbotchi/webfinger"
)

var (
	version  string
	revision string
)

func printWakeupMessage() {
	fmt.Println("==========================================================================================")
	fmt.Print("\n")
	fmt.Println("Developed by MrBotchi team and silverscat_3")
	fmt.Println(" _ __ ___  _ __")
	fmt.Println("| '_ ` _ \\| '__|")
	fmt.Println("| | | | | | |")
	fmt.Println("|_| |_| |_|_|          _       _     _")
	fmt.Println("          | |         | |     | |   (_)")
	fmt.Println("          | |__   ___ | |_ ___| |__  _")
	fmt.Println("          | '_ \\ / _ \\| __/ __| '_ \\| |")
	fmt.Println("          | |_) | (_) | || (__| | | | |")
	fmt.Println("          |_.__/ \\___/ \\__\\___|_| |_|_|")
	fmt.Println("Version:", version, ",", "Revision:", revision)
	fmt.Print("\n")
	fmt.Println("==========================================================================================")
}

func handlerFactory(app *app.App, db *sqlx.DB) map[string]handler.HandlerIf {
	var results map[string]handler.HandlerIf = map[string]handler.HandlerIf{
		// 小ネタ
		"/schwimmwagen": &api.SchwimmwagenHandler{HTTPHandler: handler.NewHandler(app)},

		// APIエンドポイント
		"/users":              &users.UsersHandler{HTTPHandler: handler.NewHandler(app), UserModel: models.NewUserModel(db)},
		"/users/{name}/token": &users.TokenHandler{HTTPHandler: handler.NewHandler(app), UserModel: models.NewUserModel(db)},

		// Activitypub
		"/":          &activitypub.ActorHandler{HTTPHandler: handler.NewHandler(app)},
		"/inbox":     &activitypub.InboxHandler{HTTPHandler: handler.NewHandler(app)},
		"/outbox":    &activitypub.OutboxHandler{HTTPHandler: handler.NewHandler(app)},
		"/publickey": &activitypub.PublickeyHandler{HTTPHandler: handler.NewHandler(app)},
	}

	return results
}

func main() {
	printWakeupMessage()

	app := app.NewApp()

	hostMeta := webfinger.HostMetaHandler{Host: app.Config.Host}
	app.Router.Get("/.well-known/host-meta", handler.HandlerFunc(hostMeta.Get).ServeHTTP)

	webfinger := wf.Default(webfinger.WebfingerResolver{UserName: app.Config.Account.Name, Host: app.Config.Host})
	webfinger.NoTLSHandler = nil
	app.Router.Get(wf.WebFingerPath, http.HandlerFunc(webfinger.Webfinger))

	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable", app.Config.DB.Host, app.Config.DB.Port, app.Config.DB.User, app.Config.DB.Password, app.Config.DB.DBname))
	if nil != err {
		log.Fatalln(err)
	}

	hs := handlerFactory(app, db)
	for endpoint, h := range hs {
		app.Router.Get(endpoint, handler.HandlerFunc(h.Get).ServeHTTP)
		app.Router.Post(endpoint, handler.HandlerFunc(h.Post).ServeHTTP)
		app.Router.Put(endpoint, handler.HandlerFunc(h.Put).ServeHTTP)
		app.Router.Delete(endpoint, handler.HandlerFunc(h.Delete).ServeHTTP)
	}

	fmt.Println("I'm HTTP listen on :" + strconv.Itoa(app.Config.Port) + ". Have a nice day!")
	go log.Fatalln(http.ListenAndServe(":"+strconv.Itoa(app.Config.Port), app.Router))
	app.ActivityDispatcher.Start()
}
