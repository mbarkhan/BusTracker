package database

import (
	"githubmbarkhanBusTracker/auth/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	config.Load()
	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {

		return err
	}
	DB = db
	return nil
}
