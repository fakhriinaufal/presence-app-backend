package users

import (
	"context"
	"time"
)

type Domain struct {
	Id int
	DepartmentId int
	Name string
	Email string
	Password string
	Dob string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	Store(ctx context.Context, domain *Domain) (Domain, error)
	GetAll() ([]Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, domain *Domain, id int) (Domain, error)
	Delete(ctx context.Context, id int) error
}

type Repository interface {
	Store(ctx context.Context, domain *Domain) (Domain, error)
	GetAll() ([]Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, domain *Domain) (Domain, error)
	Delete(ctx context.Context, id int) error
}
