package main

import (
	"github.com/da-n/banking-lib/logger"
	"github.com/da-n/banking/app"
)

func main() {
	logger.Info("Starting application...")
	app.Start()
}
