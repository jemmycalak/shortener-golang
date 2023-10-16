package configs

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"jemmy-sapta/model"
)

func InitialDatabase(config AppConfig) *gorm.DB {
	var dns = "postgres://" + config.DATABASE_USERNAME + ":" + config.DATABASE_PASSWORD + "@" + config.DATABASE_ADDRESS + "/" + config.DATABASE_NAME
	db, errorDatabaseConnection := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if errorDatabaseConnection != nil {
		panic(errorDatabaseConnection)
	}

	errorDatabaseConnection = db.AutoMigrate(
		&model.User{},
		&model.Shortener{},
		&model.StatisticShortener{},
	)
	if errorDatabaseConnection != nil {
		panic("Error Migration")
	}

	return db
}
