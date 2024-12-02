package app

import (
	"fmt"
	"log"
	"net"
	"order-service/externals"
	"order-service/internal/config/database"
	"order-service/internal/handlers"
	"order-service/internal/protobuff"
	"order-service/internal/repositories"
	"order-service/internal/services"
	"os"

	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Run() {
	var wg sync.WaitGroup

	// init database
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	// init repository
	orderRepo := repositories.NewOrderRepository(db)

	// setup grpc client connection
	productClientConn, err := grpc.Dial(os.Getenv("PRODUCT_SERVICE_URL"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Product Service: %v", err)
	}
	defer productClientConn.Close()

	productClient := externals.NewProductServiceClient(productClientConn)

	// init service
	orderSvc := services.NewOrderService(orderRepo, productClient)

	// init grpc server
	grpcServer := grpc.NewServer()

	// init handler
	orderHandler := handlers.NewProductHandler(orderSvc)
	
	// Register the handler object with the gRPC server
	protobuff.RegisterOrderServiceServer(grpcServer, orderHandler)

	// Run the gRPC server
	port := os.Getenv("GRPC_PORT")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}

	fmt.Printf("gRPC server is running on port %s\n", port)

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := grpcServer.Serve(lis); err != nil {
			panic(err)
		}
	}()

	wg.Wait()
}
