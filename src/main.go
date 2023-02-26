package main

import (
	"github.com/ffmoyano/gofer/env"
	"github.com/ffmoyano/gofer/logger"
	"os"
	"recipes/src/database"
	"recipes/src/server"
)

func init() {
	env.Read(".env")
	logger.OpenLogs("/")
	database.Open()
}

func main() {

	defer logger.CloseLogs()

	app := server.Get()

	logger.Info("Opening api at port: %s", os.Getenv("port"))
	err := app.ListenAndServe()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

}
