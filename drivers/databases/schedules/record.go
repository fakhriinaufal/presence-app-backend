package schedules

import (
	"gorm.io/gorm"
	"presence-app-backend/business/schedules"
	"time"
)

type Schedule struct {
	Id           int
	DepartmentId int
	InTime string
	OutTime string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (s *Schedule) ToDomain() schedules.Domain {
	return schedules.Domain{
		Id:           s.Id,
		DepartmentId: s.DepartmentId,
		InTime:       s.InTime,
		OutTime:      s.OutTime,
		CreatedAt:    s.CreatedAt,
		UpdatedAt:    s.UpdatedAt,
	}
}

func FromDomain(domain *schedules.Domain) *Schedule {
	return &Schedule{
		Id:           domain.Id,
		DepartmentId: domain.DepartmentId,
		InTime:       domain.InTime,
		OutTime:      domain.OutTime,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}

