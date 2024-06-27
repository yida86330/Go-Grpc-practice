package grpc_service

import (
	"context"
	pb "go_grpc_practice/pkg/proto"
)

type Server struct {
	pb.UnimplementedAccountServiceServer
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	// 模擬登錄邏輯
	if req.Username == "test" && req.Password == "password" {
		return &pb.LoginResponse{Message: "Login successful"}, nil
	}
	return &pb.LoginResponse{Message: "Invalid username or password"}, nil
}

func (s *Server) PostComment(ctx context.Context, req *pb.CommentRequest) (*pb.CommentResponse, error) {
	// 模擬發布留言邏輯
	return &pb.CommentResponse{Message: "Comment posted successfully"}, nil
}

func (s *Server) GetComments(ctx context.Context, req *pb.CommentsRequest) (*pb.CommentsResponse, error) {
	// 模擬獲取留言記錄邏輯
	comments := []string{"Comment 1", "Comment 2", "Comment 3"}
	return &pb.CommentsResponse{Comments: comments}, nil
}
