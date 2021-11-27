package presences

import (
	"gorm.io/gorm"
	"presence-app-backend/business/presences"
	"time"
)

type Presence struct {
	Id         int
	UserId     int
	ScheduleId int
	Type       string
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (p *Presence) ToDomain() presences.Domain {
	return presences.Domain{
		Id:         p.Id,
		UserId:     p.UserId,
		ScheduleId: p.ScheduleId,
		Type:       p.Type,
		Status:     p.Status,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}
}

func FromDomain(domain *presences.Domain) Presence {
	return Presence{
		Id:         domain.Id,
		UserId:     domain.UserId,
		ScheduleId: domain.ScheduleId,
		Type:       domain.Type,
		Status:     domain.Status,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}
