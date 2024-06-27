package controllers

import (
	"context"
	"log"
	"net/http"

	pb "go_grpc_practice/pkg/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var grpcClient pb.AccountServiceClient

func InitGRPCClient(conn *grpc.ClientConn) {
	grpcClient = pb.NewAccountServiceClient(conn)
}

type LoginRequest struct {
	Username string `json:"username" example:"user1"`
	Password string `json:"password" example:"password"`
}

// LoginResponse
// swagger:model LoginResponse
type LoginResponse struct {
	Token string `json:"token" example:"abc123"`
}

// @Summary 登入
// @Description 登入
// @Tags 使用者
// @Accept  json
// @Produce  json
// @Param username body LoginRequest true "login"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} LoginResponse
// @Router /login [post]
func Login(c *gin.Context) {
	var req pb.LoginRequest

	log.SetPrefix("gin api: ")
	log.Println("Starting login...")
	log.Println(&req)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := grpcClient.Login(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// CommentsRequest
// swagger:model CommentsRequest
type CommentsRequest struct{}

// CommentsResponse
// swagger:model CommentsResponse
type CommentsResponse struct {
	Comments []Comment `json:"comments"`
}

// Comment
// swagger:model Comment
type Comment struct {
	ID        int    `json:"id" example:"1"`
	Message   string `json:"message" example:"Hello"`
	Author    string `json:"author" example:"John"`
	CreatedAt string `json:"created_at" example:"2022-01-01T12:00:00Z"`
}

// @Summary 獲取留言
// @Description 獲取留言
// @Tags 留言
// @Produce  json
// @Success 200 {object} CommentsResponse
// @Router /comments [get]
func GetComments(c *gin.Context) {
	req := pb.CommentsRequest{}
	res, err := grpcClient.GetComments(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// CommentRequest
// swagger:model CommentRequest
type CommentRequest struct {
	Message string `json:"message" example:"Hello"`
}

// CommentResponse
// swagger:model CommentResponse
type CommentResponse struct {
	Message string `json:"message" example:"1"`
}

// @Summary 發布留言
// @Description 發布留言
// @Tags 留言
// @Accept  json
// @Produce  json
// @Param comment body CommentRequest true "留言内容"
// @Success 200 {object} CommentResponse
// @Router /comments [post]
func PostComment(c *gin.Context) {
	var req pb.CommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := grpcClient.PostComment(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
