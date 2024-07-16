package main

import "github.com/bugsmachine/go-logger/logger"

func main() {
	var log = logger.Logger{}
	log.Print("ERROR", "This is an error message")
}
