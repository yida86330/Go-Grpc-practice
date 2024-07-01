package models

type Comment struct {
	ID      uint32 `gorm:"primaryKey;autoIncrement"`
	Message string `json:"message"`
}
