package responses

import (
	"presence-app-backend/business/departments"
	"time"
)

type DepartmentResponse struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func FromDomain(domain departments.Domain) DepartmentResponse {
	return DepartmentResponse{
		Id:          domain.ID,
		Name:        domain.Name,
		Description: domain.Description,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
