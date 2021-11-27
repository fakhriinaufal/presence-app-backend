package users

import (
	"gorm.io/gorm"
	"presence-app-backend/business/users"
	"time"
)

type User struct {
	Id           int
	DepartmentId int
	Name         string
	Email        string
	Password     string
	Dob          string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (u *User) ToDomain() users.Domain {
	return users.Domain{
		Id:           u.Id,
		DepartmentId: u.DepartmentId,
		Name:         u.Name,
		Email:        u.Email,
		Password:     u.Password,
		Dob:          u.Dob,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}

func FromDomain(u *users.Domain) User {
	return User{
		Id:           u.Id,
		DepartmentId: u.DepartmentId,
		Name:         u.Name,
		Email:        u.Email,
		Password:     u.Password,
		Dob:          u.Dob,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}

func ToArrayOfDomain(u *[]User) []users.Domain {
	var result []users.Domain

	for _, val := range *u {
		result = append(result, val.ToDomain())
	}

	return result
}
