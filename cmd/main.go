package main

import (
	"github.com/santos95mat/go-docker-learning/internal/app"
	"github.com/santos95mat/go-docker-learning/internal/database"
)

func init() {
	database.Connect()
}

func main() {
	app.Run()
}
