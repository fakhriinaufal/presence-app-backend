package schedules

import (
	"context"
	"time"
)

type Domain struct {
	Id           int
	DepartmentId int
	InTime string
	OutTime string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	Store(ctx context.Context, domain *Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id int) (Domain ,error)
	Update(ctx context.Context, domain *Domain) (Domain, error)
}

type Repository interface {
	Store(ctx context.Context, domain *Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id int) (Domain ,error)
	Update(ctx context.Context, domain *Domain) (Domain, error)
}