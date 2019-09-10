package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mrbotchi-team/mrbotchi/app"
	"github.com/mrbotchi-team/mrbotchi/handlers"
)

var (
	version  string
	revision string
)

func PrintWakeupMessage() {
	fmt.Println("==========================================================================================")
	fmt.Print("\n")
	fmt.Println("Developed by silverscat_3")
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

func main() {
	PrintWakeupMessage()

	app := app.NewApp()

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
