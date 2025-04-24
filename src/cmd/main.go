package main

import (
	"golang-project/src/api"
	"golang-project/src/config"

	"golang-project/src/data/cache"
	"golang-project/src/data/db"
	"golang-project/src/pkg/logging"
)

func main() {
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)
	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}
	err = db.InitDb(cfg)

	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}

	defer db.CloseDb()
	logger.Info(logging.Postgres, logging.Startup, "log init", nil)
	api.InitServer()
}
