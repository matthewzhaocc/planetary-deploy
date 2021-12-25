package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/matthewzhaocc/planetary-deploy/internal"
	"github.com/matthewzhaocc/planetary-deploy/internal/config"
	"github.com/matthewzhaocc/planetary-deploy/internal/log"
	"github.com/matthewzhaocc/planetary-deploy/internal/sysinit"
	"github.com/matthewzhaocc/planetary-deploy/internal/types"
	"gorm.io/gorm"
)

var serverConfig *types.ServerConfig
var db *gorm.DB
var err error

func init() {
	path := flag.String("config.path", "config.json", "the configuration path")
	flag.Parse()

	log.Setup()
	serverConfig, err = config.ParseConfig(*path)
	if err != nil {
		fmt.Println(err)
	}

	db, err = sysinit.GetDB(*serverConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	var server = types.Server{DB: db}
	// Start API server
	http.ListenAndServe(":"+os.Getenv("PORT"), internal.GetMux(server))
}
