package responses

import (
	"presence-app-backend/business/presences"
	"time"
)

type Presence struct {
	Id         int       `json:"id"`
	UserId     int       `json:"user_id"`
	ScheduleId int       `json:"schedule_id"`
	Type       string    `json:"type"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func FromDomain(domain presences.Domain) Presence {
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
