package grpc_server

import (
	"go_grpc_practice/internal/comment"
	pb "go_grpc_practice/internal/proto"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func StartServer() {
	grpc_port := os.Getenv("GRPC_PORT")
	// 初始化 gRPC server
	listen, err := net.Listen("tcp", grpc_port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCommentServiceServer(s, &comment.Service{})
	go func() {
		if err := s.Serve(listen); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

}
