package main

import (
	"net/http"
	"os"

	"github.com/matthewzhaocc/planetary-deploy/internal"
)

func init() {
	internal.SetupLogger()
}

func main() {
	internal.LogMessage("hello")

	// Start API server
	http.ListenAndServe(":"+os.Getenv("PORT"), internal.GetMux())
}
