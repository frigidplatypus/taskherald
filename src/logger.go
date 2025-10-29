package main

import (
	"log"
	"os"
)

var logger *log.Logger

func InitLogger() {
	logger = log.New(os.Stdout, "[TaskHerald] ", log.LstdFlags)
}

func LogInfo(msg string) {
	logger.Println("INFO:", msg)
}

func LogError(err error) {
	logger.Println("ERROR:", err)
}