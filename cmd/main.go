package main

import (
	"log"
	"net"

	"go_grpc_practice/pkg/controllers"
	docs "go_grpc_practice/pkg/docs"
	"go_grpc_practice/pkg/grpc_service"
	pb "go_grpc_practice/pkg/proto"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAccountServiceServer(s, &grpc_service.Server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// gRPC client
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	controllers.InitGRPCClient(conn)

	// Gin server
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := router.Group("/api/v1")
	{
		v1.POST("/login", controllers.Login)
		v1.GET("/comments", controllers.GetComments)
		v1.POST("/comments", controllers.PostComment)
	}

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":8080")
}
