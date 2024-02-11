// cmd/server/main.go
package main

import (
	"log"
	"net"
	"os"

	"github.com/kentyisapen/simple-drive/internal/application/service"
	"github.com/kentyisapen/simple-drive/internal/application/usecase"
	"github.com/kentyisapen/simple-drive/internal/interfraces/grpc/server"
	"github.com/kentyisapen/simple-drive/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "8021"
	}
	address := ":" + port

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// 依存性を解決
	healthCheckUsecase := usecase.NewHealthCheckUsecase()
	healthCheckService := service.NewHealthCheckService(healthCheckUsecase)

	// 他のサービスやユースケースのインスタンスを作成

	// FileManagerServerのインスタンスを作成し、すべての依存性を注入
	fileManagerServer := server.NewFileManagerServer(healthCheckService /*, 他のサービスやユースケースのインスタンス */)
	pb.RegisterFileManagerServer(grpcServer, fileManagerServer)

	if os.Getenv("ENV") != "PROD" {
		reflection.Register(grpcServer)
	}

	log.Println("Server listening at", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
