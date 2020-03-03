package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/silverscat-3/hostmeta"
	hmh "github.com/silverscat-3/hostmeta/handlers"
	wf "github.com/writeas/go-webfinger"
	"gopkg.in/spacemonkeygo/httpsig.v0"

	"github.com/mrbotchi-team/mrbotchi/app"
	"github.com/mrbotchi-team/mrbotchi/handler"
	"github.com/mrbotchi-team/mrbotchi/handler/activitypub"
	"github.com/mrbotchi-team/mrbotchi/handler/api"
	"github.com/mrbotchi-team/mrbotchi/handler/api/instance"
	"github.com/mrbotchi-team/mrbotchi/handler/api/users"
	"github.com/mrbotchi-team/mrbotchi/models"
	"github.com/mrbotchi-team/mrbotchi/webfinger"
)

var (
	version  string = "Schwimmwagen"
	revision string = "Typ 166"
)

func printWakeupMessage() {
	fmt.Println("==========================================================================================")
	fmt.Print("\n")
	fmt.Println("Developed by MrBotchi team and silverscat_3")

	fmt.Println("  _ __ ___   _ __                   ")
	fmt.Println(" | '_ ` _ \\ | '__|                  ")
	fmt.Println(" | | | | | || | _                   ")
	fmt.Println(" |_|_|_| |_||_|(_)        _      _  ")
	fmt.Println(" |  _ \\        | |       | |    (_) ")
	fmt.Println(" | |_) |  ___  | |_  ___ | |__   _  ")
	fmt.Println(" |  _ <  / _ \\ | __|/ __|| '_ \\ | | ")
	fmt.Println(" | |_) || (_) || |_| (__ | | | || | ")
	fmt.Println(" |____/  \\___/  \\__|\\___||_| |_||_| ")

	fmt.Println("Version:", version, ",", "Revision:", revision)
	fmt.Print("\n")
	fmt.Println("==========================================================================================")
}

func handlerFactory(app *app.App, db *sqlx.DB, signer *httpsig.Signer) map[string]handler.HandlerIf {
	var results map[string]handler.HandlerIf = map[string]handler.HandlerIf{
		// 小ネタ
		"/schwimmwagen": &api.SchwimmwagenHandler{HTTPHandler: handler.NewHandler(app, signer)},

		// APIエンドポイント
		"/instance/specs":     &instance.SpecsHandler{HTTPHandler: handler.NewHandler(app, signer)},
		"/users":              &users.UsersHandler{HTTPHandler: handler.NewHandler(app, signer), UserModel: models.NewUserModel(db)},
		"/users/{name}/token": &users.TokenHandler{HTTPHandler: handler.NewHandler(app, signer), UserModel: models.NewUserModel(db)},

		// Activitypub
		"/":          &activitypub.ActorHandler{HTTPHandler: handler.NewHandler(app, signer)},
		"/inbox":     &activitypub.InboxHandler{HTTPHandler: handler.NewHandler(app, signer)},
		"/outbox":    &activitypub.OutboxHandler{HTTPHandler: handler.NewHandler(app, signer)},
		"/publickey": &activitypub.PublickeyHandler{HTTPHandler: handler.NewHandler(app, signer)},
	}

	return results
}

func main() {
	printWakeupMessage()

	app := app.NewApp()

	startHTTP(app)
}

func startHTTP(app *app.App) {
	links := []*hostmeta.Link{
		hostmeta.NewLink("lrdd", "application/xrd+xml", "", fmt.Sprintf("https://%s/.well-known/webfinger?resource={uri}", app.Config.Host)),
	}
	hostMeta := hmh.HostMetaHandler{Links: links}
	hostMetaJSON := hmh.HostMetaJSONHandler{Links: links}

	app.Router.Get(hostmeta.HostMetaPath, http.HandlerFunc(hostMeta.ServeHTTP))
	app.Router.Get(hostmeta.HostMetaJSONPath, http.HandlerFunc(hostMetaJSON.ServeHTTP))

	webfinger := wf.Default(webfinger.WebfingerResolver{UserName: app.Config.Account.Name, Host: app.Config.Host})
	webfinger.NoTLSHandler = nil
	app.Router.Get(wf.WebFingerPath, http.HandlerFunc(webfinger.Webfinger))

	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable", app.Config.DB.Host, app.Config.DB.Port, app.Config.DB.User, app.Config.DB.Password, app.Config.DB.DBname))
	if nil != err {
		log.Fatalln(err)
	}

	signer := httpsig.NewSigner("", app.Config.Account.PrivateKey, httpsig.RSASHA256, nil)

	hs := handlerFactory(app, db, signer)
	for endpoint, h := range hs {
		app.Router.Get(endpoint, handler.HandlerFunc(h.Get).ServeHTTP)
		app.Router.Post(endpoint, handler.HandlerFunc(h.Post).ServeHTTP)
		app.Router.Put(endpoint, handler.HandlerFunc(h.Put).ServeHTTP)
		app.Router.Delete(endpoint, handler.HandlerFunc(h.Delete).ServeHTTP)
	}

	fmt.Println("I'm HTTP listen on :" + strconv.Itoa(app.Config.Port) + ". Have a nice day!")
	log.Fatalln(http.ListenAndServe(":"+strconv.Itoa(app.Config.Port), app.Router))
}
