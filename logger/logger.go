package logger

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"
)

type Logger struct {
	Mode string
}

var Log = Logger{
	Mode: "debug",
}

func Default() Logger {
	return Logger{
		Mode: "debug",
	}
}

func (l *Logger) Release() {
	l.Mode = "release"
}

func (l *Logger) Debug() {
	l.Mode = "debug"
}

func checkMode(l Logger) bool {
	if l.Mode == "debug" {
		return true
	} else {
		return false
	}
}

func (l *Logger) Info(message string, variables ...interface{}) {
	if l.Mode == "debug" {
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

		if len(variables) > 0 {
			message = fmt.Sprintf(message, variables...)
		}

		fmt.Printf("%s %s%s\033[0m: %s (%s %d)\n", currentTime, colorCode, logType, message, shortFile, line)
	}
}

func (l *Logger) Print(logType string, message string, variables ...interface{}) {
	if l.Mode == "debug" {
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

		if len(variables) > 0 {
			message = fmt.Sprintf(message, variables...)
		}

		fmt.Printf("%s %s%s\033[0m: %s (%s %d)\n", currentTime, colorCode, logType, message, shortFile, line)
	}
}

func (l *Logger) Warning(message string, variables ...interface{}) {
	if l.Mode == "debug" {
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

		if len(variables) > 0 {
			message = fmt.Sprintf(message, variables...)
		}

		fmt.Printf("%s %s%s\033[0m: %s (%s %d)\n", currentTime, colorCode, logType, message, shortFile, line)
	}
}

func (l *Logger) Error(message string, variables ...interface{}) {
	if l.Mode == "debug" {
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

		if len(variables) > 0 {
			message = fmt.Sprintf(message, variables...)
		}

		fmt.Printf("%s %s%s\033[0m: %s (%s %d)\n", currentTime, colorCode, logType, message, shortFile, line)
	}
}
