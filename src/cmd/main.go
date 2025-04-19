package main

import (
	"golang-project/src/api"
	"golang-project/src/config"
	"golang-project/src/data/cache"
)

func main() {
	cfg :=config.GetConfig()
	cache.InitRedis(cfg)
	defer cache.CloseRedis()
	api.InitServer()
}
