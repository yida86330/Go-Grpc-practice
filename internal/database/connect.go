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
		DB: MysqlDB{},
	}

	connection, conErr := migrations.DB.GetConnection()

	if conErr != nil {
		log.Println("Connection Error Migrations: ", conErr.Error())
	}

	err := connection.AutoMigrate(&models.Comment{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}
