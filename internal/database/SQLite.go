package database

import (
	"log"

	db "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteDB struct {
}

func (sqlite SQLiteDB) GetConnection() (*gorm.DB, error) {
	conn, conErr := gorm.Open(db.Open("data.db?_journal_mode=WAL&_cache=shared"), &gorm.Config{})

	if conErr != nil {
		log.Fatal("error direct db: ", conErr)
		return nil, conErr
	}

	return conn, nil
}
