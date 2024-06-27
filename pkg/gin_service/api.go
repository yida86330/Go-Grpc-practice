package gin_service

import (
	"context"
	"net/http"

	pb "go_grpc_practice/internal/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var grpcClient pb.CommentServiceClient

func InitGRPCClient(conn *grpc.ClientConn) {
	grpcClient = pb.NewCommentServiceClient(conn)
}

type Comment struct {
	Id      string `json:"id" example:"1"`
	Message string `json:"message" example:"test msg"`
}

type httpResponse struct {
	Response string `json:"response"`
}

// @Summary 新增留言
// @Description 新增留言
// @Tags 留言
// @Accept  json
// @Produce  json
// @Param comment body Comment true "Create comment"
// @Success 200 {object} httpResponse
// @Failure 400 {object} httpResponse
// @Router /comment [post]
func PostComment(c *gin.Context) {
	var req pb.Comment

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := grpcClient.Create(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary 獲取留言
// @Description 獲取留言
// @Tags 留言
// @Accept   json
// @Produce  json
// @Param    id path string  true "comment ID"
// @Success 200 {object} httpResponse
// @Failure	400 {object}  httpResponse
// @Router /comment/{id} [get]
func GetComment(c *gin.Context) {
	id := c.Param("id")

	req := &pb.GetCommentRequest{
		Id: id,
	}

	res, err := grpcClient.Get(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary 刪除留言
// @Description 刪除留言
// @Tags 留言
// @Accept  json
// @Produce  json
// @Param id  path string true "comment ID"
// @Success 200 {object} httpResponse
// @Failure	 400 {object}  httpResponse
// @Router /comment/{id} [delete]
func DeleteComment(c *gin.Context) {
	id := c.Param("id")

	req := &pb.DeleteCommentRequest{
		Id: id,
	}

	res, err := grpcClient.Delete(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
