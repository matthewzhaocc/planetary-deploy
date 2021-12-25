package main

import (
	"net/http"
	"os"

	"github.com/matthewzhaocc/planetary-deploy/internal"
	"github.com/matthewzhaocc/planetary-deploy/internal/log"
)

func init() {
	log.Setup()
}

func main() {

	// Start API server
	http.ListenAndServe(":"+os.Getenv("PORT"), internal.GetMux())
}
