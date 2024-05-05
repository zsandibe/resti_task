package main

import (
	"github.com/zsandibe/resti_task/internal/app"
	logger "github.com/zsandibe/resti_task/pkg"
)

func main() {
	if err := app.Start(); err != nil {
		logger.Error(err)
		return
	}
}
