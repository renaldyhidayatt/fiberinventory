package config

import (
	"fiberinventory/models"
	"fiberinventory/pkg"
	"fmt"

	logger "github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	// const (
	// 	host     = "localhost"
	// 	port     = 5432
	// 	user     = "postgres"
	// 	password = ""
	// 	dbname   = "invengo"
	// )

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		pkg.GodotEnv("DB_HOST"),
		pkg.GodotEnv("DB_PORT"),
		pkg.GodotEnv("DB_USER"),
		pkg.GodotEnv("DB_PASS"),
		pkg.GodotEnv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		defer logger.Info("Database connection failed")
		logger.Fatal(err)
		return nil
	}

	err = db.AutoMigrate(&models.ModelCategory{}, &models.ModelUser{}, &models.ModelSupplier{}, &models.ModelCustomer{})
	if err != nil {
		defer logger.Info("Database connection failed")
		logger.Fatal(err)

		return nil
	}

	return db
}
