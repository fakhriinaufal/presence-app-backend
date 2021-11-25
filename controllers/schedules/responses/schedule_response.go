package responses

import (
	"presence-app-backend/business/schedules"
	"time"
)

type Schedule struct {
	Id           int `json:"id"`
	DepartmentId int `json:"department_id"`
	InTime string `json:"in_time"`
	OutTime string `json:"out_time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain schedules.Domain) Schedule {
	return Schedule{
		Id:           domain.Id,
		DepartmentId: domain.DepartmentId,
		InTime:       domain.InTime,
		OutTime:      domain.OutTime,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}
