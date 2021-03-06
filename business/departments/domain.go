package departments

import (
	"context"
	"time"
)

type Domain struct {
	ID          int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	Store(ctx context.Context, department *Domain) (Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, department *Domain) (Domain, error)
	Delete(ctx context.Context, id int) error
}

type Repository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	Store(ctx context.Context, department *Domain) (Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, department *Domain) (Domain, error)
	Delete(ctx context.Context, id int) error
}
