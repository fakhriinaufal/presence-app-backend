package requests

import (
	"presence-app-backend/business/departments"
)

type DepartmentStore struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func (dept *DepartmentStore) ToDomain() *departments.Domain {
	return &departments.Domain{
		Name:        dept.Name,
		Description: dept.Description,
	}
}
