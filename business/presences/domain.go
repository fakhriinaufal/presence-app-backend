package presences

import (
	"context"
	"time"
)

type Domain struct {
	Id         int
	UserId     int
	ScheduleId int
	Type       string
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Usecase interface {
	Store(ctx context.Context, domain *Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, domain *Domain) (Domain, error)
	Delete(ctx context.Context, id int) error
}

type Repository interface {
	Store(ctx context.Context, domain *Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, domain *Domain) (Domain, error)
	Delete(ctx context.Context, id int) error
}
