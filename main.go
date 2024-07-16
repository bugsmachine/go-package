package main

import "github.com/bugsmachine/go-logger/logger"

func main() {

	//test

	var log = logger.Default()
	log.Print("ERROR", "This is an error message")
}
