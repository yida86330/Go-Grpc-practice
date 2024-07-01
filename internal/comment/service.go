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

func (s *Service) Create(ctx context.Context, comment *pb.Comment) (*pb.CreateCommentResponse, error) {
	id := comment.Id

	newComment := models.Comment{
		ID:      id,
		Message: comment.Message,
	}

	status := dataComment.Create(newComment)

	return &pb.CreateCommentResponse{Id: newComment.ID, Response: status}, nil
}

func (s *Service) Update(ctx context.Context, comment *pb.UpdateCommentRequest) (*pb.UpdateCommentResponse, error) {
	id := comment.Id
	findComment, err := dataComment.Get(id)

	if err != nil {
		return nil, err
	}

	if findComment == nil {
		return &pb.UpdateCommentResponse{
			Id:     id,
			Status: "not found",
		}, nil
	}

	findComment.Message = comment.Message
	dataComment.Update(findComment)

	return &pb.UpdateCommentResponse{
		Id:     uint32(findComment.ID),
		Status: "Update Action",
	}, nil
}

func (s *Service) Get(ctx context.Context, comment *pb.GetCommentRequest) (*pb.GetCommentResponse, error) {
	id := comment.Id
	response, err := dataComment.Get(id)

	if err != nil {
		return nil, err
	}

	commentRes := &pb.Comment{
		Id:      response.ID,
		Message: response.Message,
	}

	return &pb.GetCommentResponse{Comment: commentRes}, nil

}

func (s *Service) List(ctx context.Context, in *pb.ListCommentRequest) (*pb.ListCommentResponse, error) {
	comments, err := dataComment.List()

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

func (s *Service) Delete(ctx context.Context, comment *pb.DeleteCommentRequest) (*pb.DeleteCommentResponse, error) {
	id := comment.Id
	findComment, err := dataComment.Get(id)

	if err != nil {
		return nil, err
	}

	if findComment == nil {
		return &pb.DeleteCommentResponse{
			Id:     id,
			Status: "not found",
		}, nil
	}

	dataComment.Delete(findComment)

	return &pb.DeleteCommentResponse{
		Id:     uint32(findComment.ID),
		Status: "Delete Action",
	}, nil
}
