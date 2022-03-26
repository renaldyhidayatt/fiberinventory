package config

import (
	"fiberinventory/models"

	logger "github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	const (
		host     = "localhost"
		port     = 5432
		user     = ""
		password = ""
		dbname   = "invengo"
	)
	dsn := "host=127.0.01 user=holyraven password= dbname=invengo port=5432 sslmode=disable"
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		defer logger.Info("Database connection failed")
		logger.Fatal(err)
		return nil
	}

	err = db.AutoMigrate(&models.ModelCategory{}, &models.ModelUser{})
	if err != nil {
		defer logger.Info("Database connection failed")
		logger.Fatal(err)

		return nil
	}

	return db
}
