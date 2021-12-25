package main

import "github.com/matthewzhaocc/planetary-deploy/internal"

func init() {
	internal.SetupLogger()
}

func main() {
	internal.LogMessage("hello")
}
