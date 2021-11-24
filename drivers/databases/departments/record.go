package departments

import (
	"gorm.io/gorm"
	"presence-app-backend/business/departments"
	"time"
)

type Department struct {
	Id          int    `gorm:"primaryKey"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (department *Department) ToDomain() departments.Domain {
	return departments.Domain{
		ID:          department.Id,
		Name:        department.Name,
		Description: department.Description,
		CreatedAt:   department.CreatedAt,
		UpdatedAt:   department.UpdatedAt,
	}
}

func FromDomain(department *departments.Domain) Department {
	return Department{
		Id:          department.ID,
		Name:        department.Name,
		Description: department.Description,
		CreatedAt: department.CreatedAt,
		UpdatedAt: department.UpdatedAt,
	}
}
