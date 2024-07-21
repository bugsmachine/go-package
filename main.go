package main

import "github.com/bugsmachine/go-package/logger"

func main() {

	//test

	var log = logger.Default()
	log.Info("This is an info message")
	log.Error("This is an error message")
	log.Warning("This is a warning message")

	logger.Error("This is an error message")
	logger.Warning("This is a warning message")

	logger.Info("This is an info message")

	logger.Print("ERROR", "This is an error message")

}
