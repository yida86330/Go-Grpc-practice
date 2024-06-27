package comment

import (
	"context"
	pb "go_grpc_practice/internal/proto"
)

type Service struct {
	pb.UnimplementedCommentServiceServer
}

func (s *Service) Create(ctx context.Context, req *pb.Comment) (*pb.CreateCommentResponse, error) {
	// 模擬新建留言邏輯

	return &pb.CreateCommentResponse{Id: "new id", Response: "success"}, nil
}

func (s *Service) Get(ctx context.Context, req *pb.GetCommentRequest) (*pb.GetCommentResponse, error) {
	// 模擬取得留言邏輯
	id := req.Id

	comment := &pb.Comment{
		Id:      id,
		Message: "Get msg",
	}
	return &pb.GetCommentResponse{Comment: comment}, nil

}

func (s *Service) Delete(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.DeleteCommentResponse, error) {
	// 模擬獲取留言記錄邏輯
	id := req.Id

	delComment := &pb.DeleteCommentResponse{
		Id:     id,
		Status: "Delete Action",
	}
	return delComment, nil
}
