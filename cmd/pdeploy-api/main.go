package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/matthewzhaocc/planetary-deploy/internal"
	"github.com/matthewzhaocc/planetary-deploy/internal/config"
	"github.com/matthewzhaocc/planetary-deploy/internal/log"
	"github.com/matthewzhaocc/planetary-deploy/internal/types"
)

var serverConfig types.ServerConfig

func init() {
	log.Setup()
	serverConfig, err := config.ParseConfig("")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	// Start API server
	http.ListenAndServe(":"+os.Getenv("PORT"), internal.GetMux())
}
