package presences

import (
	"presence-app-backend/business/presences"
	"presence-app-backend/drivers/databases/schedules"
	"presence-app-backend/drivers/databases/users"
	"time"

	"gorm.io/gorm"
)

type Presence struct {
	Id         int
	UserId     int
	User       users.User
	ScheduleId int
	Schedule   schedules.Schedule
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
		User:       p.User.ToDomain(),
		ScheduleId: p.ScheduleId,
		Schedule:   p.Schedule.ToDomain(),
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
