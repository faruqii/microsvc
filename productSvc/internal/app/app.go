package app

import (
	"fmt"
	"net"
	"os"
	"product-service/internal/config/database"
	"product-service/internal/handlers"
	"product-service/internal/protobuff/protobuff"
	"product-service/internal/repositories"
	"product-service/internal/services"
	"sync"

	"google.golang.org/grpc"
)

func Run() {
	var wg sync.WaitGroup

	// init database
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	// init repository
	productRepo := repositories.NewProductRepository(db)

	// init service
	productService := services.NewProductService(productRepo)

	// init grpc server
	grpcServer := grpc.NewServer()

	// init handler
	productHandler := handlers.NewProductHandler(productService)

	// Register the handler object with the gRPC server
	protobuff.RegisterProductServiceServer(grpcServer, productHandler)

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
