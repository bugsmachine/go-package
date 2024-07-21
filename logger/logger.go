package logger

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"
)

type Logger struct {
}

var Log = Logger{}

func Default() Logger {
	return Logger{}
}

func (l Logger) Info(message string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("Error retrieving caller info")
		return
	}

	colorCode := "\033[32m" // Green
	logType := "INFO"
	logType = "[" + logType + "]"
	shortFile := filepath.Base(file) // Get the short form of the file name
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s %s%s\033[0m: %s (%s %d)\n", currentTime, colorCode, logType, message, shortFile, line)
}

func Info(message string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("Error retrieving caller info")
		return
	}

	colorCode := "\033[32m" // Green
	logType := "INFO"
	logType = "[" + logType + "]"
	shortFile := filepath.Base(file) // Get the short form of the file name
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s %s%s\033[0m: %s (%s %d)\n", currentTime, colorCode, logType, message, shortFile, line)
}

func (l Logger) Print(logType string, message string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("Error retrieving caller info")
		return
	}

	// Determine the color based on logType
	var colorCode string
	switch logType {
	case "ERROR":
		colorCode = "\033[31m" // Red
	case "WARNING":
		colorCode = "\033[33m" // Yellow
	default:
		colorCode = "\033[0m" // Default
	}

	logType = "[" + logType + "]"
	shortFile := filepath.Base(file) // Get the short form of the file name
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s %s%s\033[0m: %s (%s %d)\n", currentTime, colorCode, logType, message, shortFile, line)
}

func Print(logType string, message string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("Error retrieving caller info")
		return
	}

	// Determine the color based on logType
	var colorCode string
	switch logType {
	case "ERROR":
		colorCode = "\033[31m" // Red
	case "WARNING":
		colorCode = "\033[33m" // Yellow
	default:
		colorCode = "\033[0m" // Default
	}

	logType = "[" + logType + "]"
	shortFile := filepath.Base(file) // Get the short form of the file name
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s %s%s\033[0m: %s (%s %d)\n", currentTime, colorCode, logType, message, shortFile, line)
}

func Warning(message string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("Error retrieving caller info")
		return
	}

	colorCode := "\033[33m" // Yellow
	logType := "WARNING"
	logType = "[" + logType + "]"
	shortFile := filepath.Base(file) // Get the short form of the file name
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s %s%s\033[0m: %s (%s %d)\n", currentTime, colorCode, logType, message, shortFile, line)
}

func (l Logger) Warning(message string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("Error retrieving caller info")
		return
	}

	colorCode := "\033[33m" // Yellow
	logType := "WARNING"
	logType = "[" + logType + "]"
	shortFile := filepath.Base(file) // Get the short form of the file name
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s %s%s\033[0m: %s (%s %d)\n", currentTime, colorCode, logType, message, shortFile, line)
}

func Error(message string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("Error retrieving caller info")
		return
	}

	colorCode := "\033[31m" // Red
	logType := "ERROR"
	logType = "[" + logType + "]"
	shortFile := filepath.Base(file) // Get the short form of the file name
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s %s%s\033[0m: %s (%s %d)\n", currentTime, colorCode, logType, message, shortFile, line)
}

func (l Logger) Error(message string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("Error retrieving caller info")
		return
	}

	colorCode := "\033[31m" // Red
	logType := "ERROR"
	logType = "[" + logType + "]"
	shortFile := filepath.Base(file) // Get the short form of the file name
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s %s%s\033[0m: %s (%s %d)\n", currentTime, colorCode, logType, message, shortFile, line)
}
