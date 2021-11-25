package users

import "time"

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
	Store(domain *Domain) (Domain, error)
}

type Repository interface {
	Store(domain *Domain) (Domain, error)
}
