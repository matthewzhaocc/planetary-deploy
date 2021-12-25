package main

import "github.com/matthewzhaocc/planetary-deploy/internal/log"

func init() {
	
}

func main() {
	log.setupLogger()
	log.logMessage("hello")
}
