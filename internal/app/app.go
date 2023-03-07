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
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"
)

// @title FiberInventory
// @version 1.0
// @description REST API for FiberInventory

// @host localhost:5000
// @BasePath /api/

// @securityDefinitions.apikey BearerAuth
// @in Header
// @name Authorization

// Run initializes whole application.
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

	app := myhandler.Init()

	go func() {
		if err := app.Listen(":" + viper.GetString("PORT")); err != nil {
			panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	_ = <-c
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	fmt.Println("Fiber was successful shutdown.")

}
