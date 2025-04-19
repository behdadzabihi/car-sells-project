package db
import (
	"fmt"
	"log"
	
	"time"
	"golang-project/src/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

var dbClient * grom.DB

func InitDb(cfg *config.Config) error{
	cnn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
			cfg.Postgres.Host,
			cfg.Postgres.Port,
			cfg.Postgres.User,
			cfg.Postgres.Password,
			cfg.Postgres.DbName,
			cfg.Postgres.SSLMode,)

			dbClient, err := gorm.Open(postgres.Open(cnn), &gorm.Config{})
			err =sqlDB.Ping()
			sqlDB, err := dbClient.DB()
			if err != nil {
				return err
			}

			sqlDB.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
			sqlDB.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
			sqlDB.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime * time.Miinute)

			log.Println("Postgres connected successfully")
			return nil
}


func GetDb() *gorm.DB {
	return dbClient
}
func CloseDb() {
	sqlDB, err := dbClient.DB()
	if err != nil {
		log.Println("Error closing database connection:", err)
		return
	}
	err = sqlDB.Close()
	if err != nil {
		log.Println("Error closing database connection:", err)
		return
	}
	log.Println("Postgres connection closed")
}