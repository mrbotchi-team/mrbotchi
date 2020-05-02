package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/comail/colog"
	"github.com/mrbotchi-team/mrbotchi/app"
	flag "github.com/spf13/pflag"
)

var (
	// バージョン情報
	version  string = "Schwimmwagen"
	revision string = "Schwimmwagen"
	codeName string = "Schwimmwagen" // シュビムワーゲン!

	// コマンドライン引数
	helpFlag       bool
	configfilePath string
)

func main() {
	flag.BoolVarP(&helpFlag, "help", "h", false, "Show help message.")
	flag.StringVarP(&configfilePath, "config-path", "c", "/etc/mrbotchi.toml", "Specifies the path to the config file.")
	flag.Parse()

	colog.Register()
	fmt.Printf("MrBotchi %s Version: %s, Revision: %s\n", codeName, version, revision)
	fmt.Println("================================================================================")

	if helpFlag {
		flag.PrintDefaults()
		os.Exit(0)
	}

	app, err := app.New(configfilePath)
	if nil != err {
		panic(err)
	}
	app.Route()

	log.Println("info: Listen HTTP on :3000. Have a nice day!")
	http.ListenAndServe(":3000", app.Router)
}
