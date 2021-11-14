package models

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	DepartmentId uint       `json:"department_id" form:"department_id"`
	InTime       *time.Time `json:"in_time" form:"in_time"`
	OutTime      *time.Time `json:"out_time" form:"out_time"`
}
