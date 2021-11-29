package users

import (
	"context"
	"errors"
	"presence-app-backend/business/departments"
	"presence-app-backend/helpers/encrpyt"
	"time"
)

type UserUsecase struct {
	Repo           Repository
	DeptRepo       departments.Repository
	contextTimeout time.Duration
}

func NewUserUsecase(repo Repository, deptRepo departments.Repository, timeout time.Duration) Usecase {
	return &UserUsecase{
		Repo:           repo,
		DeptRepo:       deptRepo,
		contextTimeout: timeout,
	}
}

func (uc UserUsecase) Store(ctx context.Context, domain *Domain) (Domain, error) {
	_, err := uc.DeptRepo.GetById(ctx, domain.DepartmentId)

	if err != nil {
		return Domain{}, errors.New("department not found")
	}

	domain.Password, err = encrpyt.Hash(domain.Password)
	if err != nil {
		return Domain{}, err
	}
	user, err := uc.Repo.Store(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (uc UserUsecase) GetAll() ([]Domain, error) {
	usersFromRepo, err := uc.Repo.GetAll()

	if err != nil {
		return []Domain{}, err
	}

	return usersFromRepo, err
}

func (uc UserUsecase) GetById(ctx context.Context, id int) (Domain, error) {
	user, err := uc.Repo.GetById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (uc UserUsecase) Update(ctx context.Context, domain *Domain, id int) (Domain, error) {
	existedUser, err := uc.Repo.GetById(ctx, id)

	if err != nil {
		return Domain{}, err
	}

	domain.Id = existedUser.Id
	domain.CreatedAt = existedUser.UpdatedAt

	result, err := uc.Repo.Update(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return result, nil

}

func (uc UserUsecase) Delete(ctx context.Context, id int) error {
	_, err := uc.Repo.GetById(ctx, id)
	if err != nil {
		return err
	}
	err = uc.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
