package database

import (
	"fmt"
	"log"
	"os"

	db "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDB struct {
}

func (mysql MysqlDB) GetConnection() (*gorm.DB, error) {
	UserName := os.Getenv("DATABASE_USER")
	Password := os.Getenv("DATABASE_PWD")
	Addr := os.Getenv("DATABASE_ADDR")
	Port := os.Getenv("DATABASE_PORT")
	Database := os.Getenv("DATABASE_NAME")

	addr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", UserName, Password, Addr, Port, Database)
	conn, conErr := gorm.Open(db.Open(addr), &gorm.Config{})

	if conErr != nil {
		log.Fatal("error direct db: ", conErr)
		return nil, conErr
	}

	// conErr = conn.AutoMigrate(&models.Comment{})
	// if conErr != nil {
	// 	log.Fatalf("failed to migrate database: %v", conErr)
	// }

	return conn, nil
}
