package app

import (
	"fiberinventory/internal/handler/api"
	"fiberinventory/internal/pb"
	"fiberinventory/pkg/dotenv"
	"fiberinventory/pkg/logger"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
func RunClient() {

	err := dotenv.Viper()

	if err != nil {
		logger.Error("error", zap.Error(err))
	}

	conn, err := grpc.Dial(viper.GetString("PORT_SERVER"), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	clientAuth := pb.NewAuthServiceClient(conn)
	clientCategory := pb.NewCategoryServiceClient(conn)
	clientCustomer := pb.NewCustomerServiceClient(conn)
	clientUser := pb.NewUserServiceClient(conn)
	clientProduct := pb.NewProductServiceClient(conn)
	clientProductKeluar := pb.NewProductKeluarServiceClient(conn)
	clientSale := pb.NewSaleServiceClient(conn)
	clientSupplier := pb.NewSupplierServiceClient(conn)
	clientProductMasuk := pb.NewProductMasukServiceClient(conn)

	myhandler := api.NewHandler(clientAuth, clientCategory, clientCustomer, clientProduct, clientProductKeluar, clientProductMasuk, clientSale, clientSupplier, clientUser)

	app := myhandler.Init()

	go func() {
		if err := app.Listen(":" + viper.GetString("PORT_CLIENT")); err != nil {
			panic(err)
		}
	}()

	fmt.Println("Server running on port " + viper.GetString("PORT_SERVER"))
	fmt.Println("Client running on port " + viper.GetString("PORT_CLIENT"))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	logger.Info("Gracefully shutting down...")
	_ = app.Shutdown()

	logger.Info("Running cleanup tasks...")

	logger.Info("Fiber was successfully shut down.")

}
