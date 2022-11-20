package dto

import "gorm.io/gorm"

type CreateToDoDto struct {
	gorm.Model
	Name        string `uri:"name" json:"name"`
	Description string `uri:"name" json:"description"`
}

type UpdateToDoDto struct {
	gorm.Model
	Name        string `uri:"name" json:"name"`
	Description string `uri:"name" json:"description"`
}
