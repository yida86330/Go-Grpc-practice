package database

import (
	"fmt"
	"log"
	"time"

	db "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	UserName     string = "root"
	Password     string = "password"
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
	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", UserName, Password, Addr, Port, Database)
	conn, conErr := gorm.Open(db.Open(addr), &gorm.Config{})

	if conErr != nil {
		log.Fatal("error direct db: ", conErr)
		return nil, conErr
	}

	db, err := conn.DB()
	if err != nil {
		log.Fatal("error direct db: ", conErr)
		return nil, conErr
	}

	db.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second)
	db.SetMaxIdleConns(MaxIdleConns)
	db.SetMaxOpenConns(MaxOpenConns)

	return conn, nil
}
