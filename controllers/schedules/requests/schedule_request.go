package requests

import "presence-app-backend/business/schedules"

type Schedule struct {
	DepartmentId int `json:"department_id"`
	InTime string `json:"in_time"`
	OutTime string `json:"out_time"`
}

func (s *Schedule) ToDomain() *schedules.Domain {
	return &schedules.Domain{
		DepartmentId: s.DepartmentId,
		InTime: s.InTime,
		OutTime: s.OutTime,
	}
}
