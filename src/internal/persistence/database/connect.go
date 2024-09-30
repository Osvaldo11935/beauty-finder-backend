package database

import (
	"gorm.io/gorm"
	"src/internal/configs"
	"gorm.io/driver/postgres"
)

func Connect() (*gorm.DB, error) {
	config, configErr := configs.LoadConfig()

	if configErr != nil {
		return nil, configErr
	}

	db, openConnectionErr := gorm.Open(postgres.Open(config.PgConnectionString))

	if openConnectionErr != nil {
		return nil, configErr
	}

	return db, nil
}
