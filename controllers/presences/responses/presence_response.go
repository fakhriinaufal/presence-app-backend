package responses

import (
	"presence-app-backend/business/presences"
	"presence-app-backend/business/schedules"
	"presence-app-backend/business/users"
	"time"
)

type UserResponse struct {
	Id           int       `json:"id"`
	DepartmentId int       `json:"department_id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Dob          string    `json:"dob"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func userFromDomain(u users.Domain) UserResponse {
	return UserResponse{
		Id:           u.Id,
		DepartmentId: u.DepartmentId,
		Name:         u.Name,
		Email:        u.Email,
		Dob:          u.Dob,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}

type Schedule struct {
	Id           int       `json:"id"`
	DepartmentId int       `json:"department_id"`
	InTime       string    `json:"in_time"`
	OutTime      string    `json:"out_time"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func scheduleFromDomain(domain schedules.Domain) Schedule {
	return Schedule{
		Id:           domain.Id,
		DepartmentId: domain.DepartmentId,
		InTime:       domain.InTime,
		OutTime:      domain.OutTime,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}

type Presence struct {
	Id         int          `json:"id"`
	UserId     int          `json:"user_id"`
	User       UserResponse `json:"user"`
	ScheduleId int          `json:"schedule_id"`
	Schedule   Schedule     `json:"schedule"`
	Type       string       `json:"type"`
	Status     string       `json:"status"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
}

func FromDomain(domain presences.Domain) Presence {
	return Presence{
		Id:         domain.Id,
		UserId:     domain.UserId,
		User:       userFromDomain(domain.User),
		ScheduleId: domain.ScheduleId,
		Schedule:   scheduleFromDomain(domain.Schedule),
		Type:       domain.Type,
		Status:     domain.Status,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}
