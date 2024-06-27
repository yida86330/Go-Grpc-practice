package main

import (
	"log"

	"go_grpc_practice/internal/grpc_server"
	"go_grpc_practice/pkg/gin_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// @title Gin-Gonic API with gRPC
// @version 1.0
// @description This is a sample server for a Gin-Gonic API with gRPC integration.
// @host localhost:8080
// @BasePath /api/v1

func main() {
	// 初始化 gRPC server
	grpc_server.StartServer()

	// gRPC client
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	gin_service.InitGRPCClient(conn)

	// Gin server
	gin_service.Start("8080")
}
