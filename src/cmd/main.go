package main

import (
	"golang-project/src/api"
	"golang-project/src/config"
	"golang-project/src/data/cache"
)

func main() {
	cfg :=config.GetConfig()
	cache.InitRedis(cfg)
	db.InitDb(cfg)
	defer cache.CloseRedis()
	defer db.CloseDb()
	api.InitServer()
}
