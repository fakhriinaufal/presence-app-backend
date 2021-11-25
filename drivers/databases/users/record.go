package users

import (
	"gorm.io/gorm"
	"presence-app-backend/business/users"
	"time"
)

type Users struct {
	Id int
	DepartmentId int
	Name string
	Email string
	Password string
	Dob string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *Users) ToDomain() users.Domain {
	return users.Domain{
		u.Id,
		u.DepartmentId,
		u.Name,
		u.Email,
		u.Password,
		u.Dob,
		u.CreatedAt,
		u.UpdatedAt,
	}
}

func FromDomain(u *users.Domain) Users {
	return Users{
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