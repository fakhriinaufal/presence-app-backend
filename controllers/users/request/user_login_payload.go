package request

import "presence-app-backend/business/users"

type UserLoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserLoginPayload) ToDomain() *users.Domain {
	return &users.Domain{
		Email:    u.Email,
		Password: u.Password,
	}
}
