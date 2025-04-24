package migration

import (
	"golang-project/src/config"
	"golang-project/src/data/db"
	"golang-project/src/data/models"
	"golang-project/src/pkg/logging"
)

var logger = logging.NewLogger(config.GetConfig())

func Up_1() {

	database := db.GetDb()

	table := []interface{}{}

	country := models.Country{}
	city := models.City{}

	if !database.Migrator().HasTable(country) {
		table = append(table, country)
	}

	if !database.Migrator().HasTable(city) {
		table = append(table, city)
	}

	database.Migrator().CreateTable(table...)
	logger.Info(logging.Postgres,logging.Migration,"tables created",nil)

}
