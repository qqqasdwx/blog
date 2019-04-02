package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/qqqasdwx/blog/config"
	"github.com/qqqasdwx/blog/http"
	"github.com/qqqasdwx/blog/models"
)

func prepare() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func init() {
	prepare()

	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	help := flag.Bool("h", false, "help")
	flag.Parse()

	handleVersion(*version)
	handleHelp(*help)
	handleConfig(*cfg)
}

func main() {
	models.InitMysql()
	models.InitCache()
	go http.Run()

	select {}
}

func handleVersion(displayVersion bool) {
	if displayVersion {
		fmt.Println(config.VERSION)
		os.Exit(0)
	}
}

func handleHelp(displayHelp bool) {
	if displayHelp {
		flag.Usage()
		os.Exit(0)
	}
}

func handleConfig(configFile string) {
	err := config.Parse(configFile)
	if err != nil {
		log.Fatalln(err)
	}
}
