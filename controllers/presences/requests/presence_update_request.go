package requests

import "presence-app-backend/business/presences"

type PresenceUpdate struct {
	UserId     int    `json:"user_id"`
	ScheduleId int    `json:"schedule_id"`
	Type       string `json:"type"`
	Status     string `json:"status"`
}

func (p *PresenceUpdate) ToDomain() *presences.Domain {
	return &presences.Domain{
		UserId:     p.UserId,
		ScheduleId: p.ScheduleId,
		Type:       p.Type,
		Status:     p.Status,
	}
}
