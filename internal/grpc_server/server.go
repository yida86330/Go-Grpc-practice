package grpc_server

import (
	"go_grpc_practice/internal/comment"
	pb "go_grpc_practice/internal/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func StartServer() {
	// 初始化 gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCommentServiceServer(s, &comment.Service{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

}
