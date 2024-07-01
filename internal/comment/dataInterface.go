package comment

import (
	"go_grpc_practice/internal/database"
	"go_grpc_practice/internal/models"
)

var dBConnection database.IConnection

type CommentInterFace interface {
	Create(comment models.Comment) string
	Update(comment *models.Comment) (*models.Comment, error)
	Get(id uint32) (*models.Comment, error)
	List() ([]models.Comment, error)
	Delete(comment *models.Comment) (*models.Comment, error)
}

type dataAccess struct{}

func init() {
	dBConnection = database.MysqlDB{}
}

func (d dataAccess) Create(comment models.Comment) string {
	connection, err := dBConnection.GetConnection()

	if err != nil {
		return "DB ERROR: " + err.Error()
	}

	connection.Create(&comment)

	return "created"
}

func (d dataAccess) Update(comment *models.Comment) (*models.Comment, error) {
	connection, err := dBConnection.GetConnection()

	if err != nil {
		return nil, err
	}

	connection.Save(&comment)

	return &models.Comment{}, nil
}

func (d dataAccess) Get(id uint32) (*models.Comment, error) {
	connection, err := dBConnection.GetConnection()

	comment := models.Comment{}

	if err != nil {
		return nil, err
	}

	connection.First(&comment, "id = ?", id)

	return &comment, nil
}

func (d dataAccess) List() ([]models.Comment, error) {
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

func (d dataAccess) Delete(comment *models.Comment) (*models.Comment, error) {
	connection, err := dBConnection.GetConnection()

	if err != nil {
		return nil, err
	}

	connection.Delete(comment)

	return &models.Comment{}, nil
}
