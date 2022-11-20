package models

import "gorm.io/gorm"

type ToDoModels struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255);not null" json:"name"`
	Description string `gorm:"type:varchar(255)" json:"descriptio" `
}
