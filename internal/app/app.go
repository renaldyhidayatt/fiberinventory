package app

import (
	"fiberinventory/internal/handler"
	"fiberinventory/internal/repository"
	"fiberinventory/internal/service"
	"fiberinventory/pkg/auth"
	"fiberinventory/pkg/database/postgres"
	"fiberinventory/pkg/dotenv"
	"fiberinventory/pkg/hash"
	"fiberinventory/pkg/logger"

	"github.com/spf13/viper"
)

func Run() {

	log := logger.New(true)
	err := dotenv.Viper()

	if err != nil {
		log.Err(err)
	}

	db, err := postgres.NewClient()

	if err != nil {
		log.Err(err)
	}

	hashier := hash.NewHashingPassword()
	repository := repository.NewRepository(db)
	token, err := auth.NewManager(viper.GetString("JWT_SECRET"))

	if err != nil {
		log.Err(err)
	}

	services := service.NewService(service.Deps{
		Repository: repository,
		Hashing:    *hashier,
		Token:      token,
	})

	myhandler := handler.NewHandler(services)

	myhandler.Init().Listen("8080")

}
