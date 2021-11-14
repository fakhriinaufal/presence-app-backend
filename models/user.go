package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	DepartmentId uint       `json:"department_id" form:"department_id"`
	Name         string     `json:"name" form:"name"`
	Email        string     `json:"email" form:"email"`
	Password     string     `json:"passsword" form:"passsword"`
	Dob          *time.Time `json:"dob" form:"dob"`
}
