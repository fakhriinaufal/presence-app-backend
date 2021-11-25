package users

import "time"

type UserUsecase struct {
	Repo Repository
	contextTimeout time.Duration
}

func NewUserUsecase(repo Repository, timeout time.Duration) Usecase {
	return &UserUsecase{
		Repo: repo,
		contextTimeout: timeout,
	}
}

func (u UserUsecase) Store(domain *Domain) (Domain, error) {
	user, err := u.Repo.Store(domain)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}
