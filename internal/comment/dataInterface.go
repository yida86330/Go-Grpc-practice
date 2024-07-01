package comment

import (
	"go_grpc_practice/internal/database"
	"go_grpc_practice/internal/models"
)

var dBConnection database.IConnection

type CommentInterFace interface {
	CreateComment(comment models.Comment) string
	GetComment(id uint32) (*models.Comment, error)
	ListComment() ([]models.Comment, error)
	DeleteComment(comment *models.Comment) (*models.Comment, error)
}

type dataAccess struct{}

func init() {
	dBConnection = database.MysqlDB{}
}

func (d dataAccess) CreateComment(comment models.Comment) string {
	connection, err := dBConnection.GetConnection()

	if err != nil {
		return "DB ERROR: " + err.Error()
	}

	connection.Create(&comment)

	return "created"
}

func (d dataAccess) GetComment(id uint32) (*models.Comment, error) {
	connection, err := dBConnection.GetConnection()

	comment := models.Comment{}

	if err != nil {
		return nil, err
	}

	connection.First(&comment, "id = ?", id)

	return &comment, nil
}

func (d dataAccess) ListComment() ([]models.Comment, error) {
	connection, err := dBConnection.GetConnection()
	var comments []models.Comment

	if err != nil {
		return nil, err
	}

	result := connection.Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}

	return comments, nil
}

func (d dataAccess) DeleteComment(comment *models.Comment) (*models.Comment, error) {
	connection, err := dBConnection.GetConnection()

	if err != nil {
		return nil, err
	}

	connection.Delete(comment)

	return &models.Comment{}, nil
}
