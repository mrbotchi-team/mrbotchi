package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

var version string

type App struct {
	router *chi.Mux
}

func NewApp() *App {
	return &App{
		router: chi.NewRouter(),
	}
}

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

	app := NewApp()

	fmt.Println("I'm HTTP listen on :3000. Have a nice day!")
	log.Fatalln(http.ListenAndServe(":3000", app.router))
}
