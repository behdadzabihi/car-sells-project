package main

import (
	"log"

	"golang-project/src/api"
	"golang-project/src/config"
	"golang-project/src/data/cache"
	"golang-project/src/data/db"
)

func main() {
	cfg :=config.GetConfig()
	var err error
	err = db.InitDb(cfg)

	if err != nil {
		log.Fatalf("error in init db %v", err)
	}
	defer cache.CloseRedis()
	defer db.CloseDb()
	api.InitServer()
}
