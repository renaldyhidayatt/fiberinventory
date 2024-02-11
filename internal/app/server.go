package app

import (
	"fiberinventory/internal/handler/gapi"
	"fiberinventory/internal/pb"
	"fiberinventory/internal/repository"
	"fiberinventory/internal/service"
	"fiberinventory/pkg/auth"
	"fiberinventory/pkg/database/postgres"
	"fiberinventory/pkg/dotenv"
	"fiberinventory/pkg/hash"
	"fiberinventory/pkg/logger"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "gRPC server port")
)

func RunServer() {

	logger.Info("Grpc Server running ...")

	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		logger.Fatal("Failed to listen", zap.Error(err))
	}

	err = dotenv.Viper()

	if err != nil {
		logger.Error("error", zap.Error(err))
	}

	db, err := postgres.NewClient()

	if err != nil {
		logger.Error("error", zap.Error(err))
	}

	hashier := hash.NewHashingPassword()
	repository := repository.NewRepository(db)
	token, err := auth.NewManager(viper.GetString("JWT_SECRET"))

	if err != nil {
		logger.Error("error", zap.Error(err))
	}

	services := service.NewService(service.Deps{
		Repository: repository,
		Hashing:    *hashier,
		Token:      token,
	})

	authGapi := gapi.NewAuthHandlerGrpcHandler(services.User)

	categoryGapi := gapi.NewCategoryHandlerGrpcHandler(services.Category)

	customerGapi := gapi.NewCustomerHandlerGrpc(services.Customer)

	productGapi := gapi.NewProductHandlerGrpcHandler(services.Product)

	userGapi := gapi.NewUserHandlerGrpcHandler(services.User)

	productKeluarGapi := gapi.NewProductKeluarHandlerGrpcHandler(services.ProductKeluar)

	saleGapi := gapi.NewSaleHandlerGrpc(services.Sale)

	supplierGapi := gapi.NewSupplierHandlerGrpc(services.Supplier)

	if err != nil {
		return
	}

	s := grpc.NewServer()

	pb.RegisterAuthServiceServer(s, authGapi)
	pb.RegisterCategoryServiceServer(s, categoryGapi)
	pb.RegisterCustomerServiceServer(s, customerGapi)
	pb.RegisterProductServiceServer(s, productGapi)
	pb.RegisterUserServiceServer(s, userGapi)
	pb.RegisterProductKeluarServiceServer(s, productKeluarGapi)
	pb.RegisterSaleServiceServer(s, saleGapi)
	pb.RegisterSupplierServiceServer(s, supplierGapi)

	fmt.Println("Server running on port " + viper.GetString("PORT_SERVER"))
	fmt.Println("Client running on port " + viper.GetString("PORT_CLIENT"))

	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}

}
