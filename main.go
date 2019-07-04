package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mr-botchi/backend/app"
	"github.com/mr-botchi/backend/handlers"
)

var version string

func main() {
	fmt.Println(" _ __ ___  _ __")
	fmt.Println("| '_ ` _ \\| '__|")
	fmt.Println("| | | | | | |")
	fmt.Println("|_| |_| |_|_|          _       _     _")
	fmt.Println("          | |         | |     | |   (_)")
	fmt.Println("          | |__   ___ | |_ ___| |__  _")
	fmt.Println("          | '_ \\ / _ \\| __/ __| '_ \\| |")
	fmt.Println("          | |_) | (_) | || (__| | | | |")
	fmt.Println("          |_.__/ \\___/ \\__\\___|_| |_|_| Backend", version)
	fmt.Println("Guten Morgen!!")

	app := app.NewApp(version)

	handlers := handlers.HandlerFactory(app)
	for endpoint, handler := range handlers {
		app.Router.Get(endpoint, handler.Get)
		app.Router.Post(endpoint, handler.Post)
		app.Router.Put(endpoint, handler.Put)
		app.Router.Delete(endpoint, handler.Delete)
	}

	fmt.Println("I'm HTTP listen on :3000. Have a nice day!")
	log.Fatalln(http.ListenAndServe(":3000", app.Router))
}
