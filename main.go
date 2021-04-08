package main

import (
	"github.com/da-n/banking/app"
	"github.com/da-n/banking/logger"
)

func main() {
	logger.Info("Starting application...")
	app.Start()
}
