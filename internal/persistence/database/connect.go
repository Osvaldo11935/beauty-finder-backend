package database

import (
	"time"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"src/internal/configs"
)

var dbInstance *gorm.DB

func Connect() (*gorm.DB, error) {
	if dbInstance != nil {
		sqlDB, err := dbInstance.DB()
		if err == nil && sqlDB.Ping() == nil {
			return dbInstance, nil
		}
	}


	config, configErr := configs.LoadConfig()
	if configErr != nil {
		return nil, configErr
	}
	db, openConnectionErr := gorm.Open(postgres.New(postgres.Config{DSN: config.PgConnectionString, PreferSimpleProtocol: true,},), &gorm.Config{})
	if openConnectionErr != nil {
		return nil, openConnectionErr
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(10 * time.Minute)

	dbInstance = db
	return dbInstance, nil
}
