package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/comail/colog"
	"github.com/mrbotchi-team/mrbotchi/app"
)

var (
	version  string = "Schwimmwagen"
	revision string = "Schwimmwagen"
	codeName string = "Schwimmwagen"
)

func main() {
	colog.Register()
	fmt.Printf("MrBotchi %s Version: %s, Revision: %s\n", codeName, version, revision)
	fmt.Println("================================================================================")

	app := app.New()
	app.Route()

	log.Println("info: Listen HTTP on :3000. Have a nice day!")
	http.ListenAndServe(":3000", app.Router)
}
