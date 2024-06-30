package database

import (
	"go_grpc_practice/internal/models"
	"log"
)

type Migrations struct {
	DB IConnection
}

func InitMigrations() {
	migrations := Migrations{
		DB: SQLiteDB{},
	}

	connection, conErr := migrations.DB.GetConnection()

	if conErr != nil {
		log.Println("Connection Error Migrations: ", conErr.Error())
	}

	connection.AutoMigrate(&models.Comment{})
}
