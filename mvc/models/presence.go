package models

import "gorm.io/gorm"

type Presence struct {
	gorm.Model
	UserId uint   `json:"user_id" form:"user_id"`
	Type   string `json:"type" form:"type"`
}
