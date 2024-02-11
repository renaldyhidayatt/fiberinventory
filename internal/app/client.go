package app

import (
	"context"
	"fiberinventory/internal/handler/api"
	"fiberinventory/internal/pb"
	"fiberinventory/pkg/dotenv"
	"fiberinventory/pkg/logger"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
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
	flag.Parse()

	err := dotenv.Viper()
	if err != nil {
		logger.Error("error", zap.Error(err))
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := grpc.DialContext(ctx, *addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("Failed to connect to server", zap.Error(err))
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
		if err := app.Listen(":5000"); err != nil {
			logger.Error("server error", zap.Error(err))
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	logger.Info("Shutting down the server...")

	cancel()

	logger.Info("Server has been shut down gracefully")
}
