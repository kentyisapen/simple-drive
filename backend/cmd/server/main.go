// cmd/server/main.go
package main

import (
	"log"
	"net"
	"os"

	"github.com/kentyisapen/simple-drive/internal/interfraces/grpc/server"
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
	server.BootstrapServices(grpcServer)

	if os.Getenv("ENV") != "PROD" {
		reflection.Register(grpcServer)
	}

	log.Println("Server listening at", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
