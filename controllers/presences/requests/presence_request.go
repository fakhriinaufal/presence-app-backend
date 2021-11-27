package requests

import "presence-app-backend/business/presences"

type Presence struct {
	UserId     int    `json:"user_id"`
	ScheduleId int    `json:"schedule_id"`
	Type       string `json:"type"`
}

func (p *Presence) ToDomain() *presences.Domain {
	return &presences.Domain{
		UserId:     p.UserId,
		ScheduleId: p.ScheduleId,
		Type:       p.Type,
	}
}
