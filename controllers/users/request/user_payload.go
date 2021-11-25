package request

import "presence-app-backend/business/users"

type UserPayload struct {
	DepartmentId int    `json:"department_id" validate:"required"`
	Name         string `json:"name" validate:"required"`
	Email        string `json:"email" validate:"required"`
	Password     string `json:"password" validate:"required"`
	Dob          string `json:"dob"`
}

func (user UserPayload) ToDomain() users.Domain {
	return users.Domain{
		DepartmentId: user.DepartmentId,
		Name:         user.Name,
		Email:        user.Email,
		Password:     user.Password,
		Dob:          user.Dob,
	}
}
