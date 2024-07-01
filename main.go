package main

import (
	"log"
	"os"

	"go_grpc_practice/internal/database"
	"go_grpc_practice/internal/grpc_server"
	"go_grpc_practice/pkg/gin_service"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// @title Gin-Gonic API with gRPC
// @version 1.0
// @description This is a sample server for a Gin-Gonic API with gRPC integration.
// @host localhost:8080
// @BasePath /api/v1

func main() {
	godotenv.Load()
	grpc_port := os.Getenv("GRPC_PORT")
	gin_port := os.Getenv("GIN_PORT")

	// 初始化 gRPC server
	grpc_server.StartServer()

	// gRPC client
	conn, err := grpc.NewClient(grpc_port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	gin_service.InitGRPCClient(conn)

	//Init Sqlite
	database.InitMigrations()

	// Gin server
	gin_service.Start(gin_port)
}
