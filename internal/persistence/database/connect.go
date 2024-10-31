package database

import (
	"time"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"src/internal/configs"
)

func Connect() (*gorm.DB, error) {

	config, configErr := configs.LoadConfig()
	if configErr != nil {
		return nil, configErr
	}

	db, openConnectionErr := gorm.Open(postgres.Open(config.PgConnectionString + "?statement_cache_mode=disabled"), &gorm.Config{})
	if openConnectionErr != nil {
		return nil, openConnectionErr
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(30 * time.Minute) 

	return db, nil
}
