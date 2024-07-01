package comment

import (
	"context"
	"go_grpc_practice/internal/models"
	pb "go_grpc_practice/internal/proto"
)

var dataComment CommentInterFace

func init() {
	dataComment = dataAccess{}
}

type Service struct {
	pb.UnimplementedCommentServiceServer
}

func (s *Service) CreateComment(ctx context.Context, comment *pb.Comment) (*pb.CreateCommentResponse, error) {
	id := comment.Id

	newComment := models.Comment{
		ID:      id,
		Message: comment.Message,
	}

	status := dataComment.CreateComment(newComment)

	return &pb.CreateCommentResponse{Id: newComment.ID, Response: status}, nil
}

func (s *Service) GetComment(ctx context.Context, comment *pb.GetCommentRequest) (*pb.GetCommentResponse, error) {
	id := comment.Id
	response, err := dataComment.GetComment(id)

	if err != nil {
		return nil, err
	}

	commentRes := &pb.Comment{
		Id:      response.ID,
		Message: response.Message,
	}

	return &pb.GetCommentResponse{Comment: commentRes}, nil

}

func (s *Service) ListComment(ctx context.Context, in *pb.ListCommentRequest) (*pb.ListCommentResponse, error) {
	comments, err := dataComment.ListComment()

	if err != nil {
		return nil, err
	}

	var commentRes []*pb.Comment
	for _, comment := range comments {
		pbComment := &pb.Comment{
			Id:      uint32(comment.ID),
			Message: comment.Message,
		}
		commentRes = append(commentRes, pbComment)
	}

	return &pb.ListCommentResponse{Comments: commentRes}, nil

}

func (s *Service) DeleteComment(ctx context.Context, comment *pb.DeleteCommentRequest) (*pb.DeleteCommentResponse, error) {
	id := comment.Id
	findComment, err := dataComment.GetComment(id)

	if err != nil {
		return nil, err
	}

	if findComment == nil {
		return &pb.DeleteCommentResponse{
			Id:     id,
			Status: "not found",
		}, nil
	}

	dataComment.DeleteComment(findComment)

	return &pb.DeleteCommentResponse{
		Id:     uint32(findComment.ID),
		Status: "Delete Action",
	}, nil
}
