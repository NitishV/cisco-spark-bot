package main

import (
	"log"
	"os"

	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	engine := getWebEngine()
	engine.Run(":" + port)
}
