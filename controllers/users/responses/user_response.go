package responses

import (
	"presence-app-backend/business/users"
	"time"
)

type UserResponse struct {
	Id int `json:"id"`
	DepartmentId int `json:"department_id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Dob string `json:"dob"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *UserResponse) FromDomain() users.Domain {
	return users.Domain{
		Id:           u.Id,
		DepartmentId: u.DepartmentId,
		Name:         u.Name,
		Email:        u.Email,
		Dob:          u.Dob,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}
