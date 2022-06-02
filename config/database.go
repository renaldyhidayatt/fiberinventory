package config

import (
	"encoding/json"
	"fiberinventory/models"
	"fiberinventory/pkg"
	"fiberinventory/schemas"
	"fmt"
	"log"
	"os"

	logger "github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Config = LoadConfiguration()

func SetupDatabase() *gorm.DB {
	// const (
	// 	host     = "localhost"
	// 	port     = 5432
	// 	user     = "postgres"
	// 	password = ""
	// 	dbname   = "invengo"
	// )

	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v",
		pkg.GodotEnv("PG_HOST"), pkg.GodotEnv("PG_USER"), pkg.GodotEnv("PG_PASS"),
		pkg.GodotEnv("PG_DBNM"), pkg.GodotEnv("PG_PORT"),
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

func LoadConfiguration() schemas.Configuration {
	var configuration schemas.Configuration
	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatal("Error reading the file")
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Fatal("can't decode config JSON: ", err)
	}
	return configuration
}
