package gin_service

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"go_grpc_practice/internal/models"
	pb "go_grpc_practice/internal/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var grpcClient pb.CommentServiceClient

func InitGRPCClient(conn *grpc.ClientConn) {
	grpcClient = pb.NewCommentServiceClient(conn)
}

type httpResponse struct {
	Response string `json:"response"`
}

// @Summary 新增留言
// @Description 新增留言
// @Tags 留言
// @Accept  json
// @Produce  json
// @Param comment body models.Comment true "Create comment"
// @Success 200 {object} httpResponse
// @Failure 400 {object} httpResponse
// @Router /comment [post]
func CreateComment(c *gin.Context) {
	commentModel := models.Comment{}
	commentProto := &pb.Comment{}

	err := c.ShouldBindJSON(&commentModel)

	commentProto.Id = uint32(commentModel.ID)
	commentProto.Message = commentModel.Message

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := grpcClient.Create(context.Background(), commentProto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary 更新留言
// @Description 更新留言
// @Tags 留言
// @Accept   json
// @Produce  json
// @Param    id path string true "comment ID"
// @Param    message body models.Comment true "Update message"
// @Success 200 {object} httpResponse
// @Failure	400 {object}  httpResponse
// @Router /comment/{id} [put]
func UpdateComment(c *gin.Context) {
	commentModel := models.Comment{}
	commentProto := &pb.UpdateCommentRequest{}

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID empty"})
		return
	}

	idUint32, _ := strconv.ParseUint(id, 10, 32)

	err := c.ShouldBindJSON(&commentModel)

	commentProto.Id = uint32(idUint32)
	commentProto.Message = commentModel.Message
	log.Println(&commentModel)
	log.Println(commentProto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := grpcClient.Update(context.Background(), commentProto)
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

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID empty"})
		return
	}

	idUint32, err := strconv.ParseUint(id, 10, 32)

	req := &pb.GetCommentRequest{
		Id: uint32(idUint32),
	}

	res, err := grpcClient.Get(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary 列出所有留言
// @Description 列出所有留言
// @Tags 留言
// @Accept   json
// @Produce  json
// @Success 200 {object} httpResponse
// @Failure	400 {object}  httpResponse
// @Router /comments/ [get]
func ListComment(c *gin.Context) {

	res, err := grpcClient.List(context.Background(), &pb.ListCommentRequest{})
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

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID empty"})
		return
	}

	idUint32, err := strconv.ParseUint(id, 10, 32)

	req := &pb.DeleteCommentRequest{
		Id: uint32(idUint32),
	}

	res, err := grpcClient.Delete(context.Background(), req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
