package models

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
}
