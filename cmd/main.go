package main

import (
	"github.com/zsandibe/resti_task/internal/app"
	logger "github.com/zsandibe/resti_task/pkg"
)

// @title Resti API
// @version 1.0
// @description This is basic server for a car service
// @host localhost:8888
// @BasePath /api/v1
func main() {
	if err := app.Start(); err != nil {
		logger.Error(err)
		return
	}
}
