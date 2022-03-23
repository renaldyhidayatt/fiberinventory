package config

import (
	"fmt"

	logger "github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	const (
		host     = "localhost"
		port     = 5432
		user     = "holyraven"
		password = ""
		dbname   = "inventory"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		defer logger.Info("Database connection failed")
		logger.Fatal(err)

		return nil
	}

	return db
}
