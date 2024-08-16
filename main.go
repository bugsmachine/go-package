package main

import (
	"fmt"
	"github.com/bugsmachine/go-package/logger"
	"strconv"
)

func main() {

	//test

	var log = logger.Default()
	//log.Release()
	fmt.Println(log.Mode)
	log.Info("This is an info message")
	a := "12138aaa"
	_, err := strconv.Atoi(a)
	log.Info("This is an info message with error: %s", err)
	log.Info("This is an info message with variables: %s, and  %d", "test", 123)
	log.Error("This is an error message")
	log.Error("This is an error message with error: %s", err)
	log.Warning("This is a warning message")
	log.Warning("This is a warning message with error: %s", err)

}
