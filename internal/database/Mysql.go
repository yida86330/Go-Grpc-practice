package database

import (
	"fmt"
	"log"

	db "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	UserName     string = "root"
	Password     string = "Zxc86330!#"
	Addr         string = "127.0.0.1"
	Port         int    = 3306
	Database     string = "test"
	MaxLifetime  int    = 10
	MaxOpenConns int    = 10
	MaxIdleConns int    = 10
)

type MysqlDB struct {
}

func (mysql MysqlDB) GetConnection() (*gorm.DB, error) {
	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", UserName, Password, Addr, Port, Database)
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
