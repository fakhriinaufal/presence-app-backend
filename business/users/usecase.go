package users

import (
	"context"
	"errors"
	"fmt"
	"presence-app-backend/app/middlewares"
	"presence-app-backend/business/departments"
	"presence-app-backend/helpers/encrpyt"
	"strings"
	"time"
)

type UserUsecase struct {
	JwtConfig      middlewares.ConfigJWT
	Repo           Repository
	DeptRepo       departments.Repository
	contextTimeout time.Duration
}

func NewUserUsecase(repo Repository, deptRepo departments.Repository, timeout time.Duration, jwtConfig middlewares.ConfigJWT) Usecase {
	return &UserUsecase{
		JwtConfig:      jwtConfig,
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

func (uc UserUsecase) Login(ctx context.Context, domain *Domain) (Domain, error) {
	var err error

	result, err := uc.Repo.GetByEmail(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	if !encrpyt.ValidateHash(domain.Password, result.Password) {
		return Domain{}, errors.New("password mismatch")
	}

	resultDomain, err := uc.DeptRepo.GetById(ctx, result.DepartmentId)
	if err != nil {
		return Domain{}, err
	}

	fmt.Println(resultDomain.Name)
	if strings.ToLower(resultDomain.Name) == "admin" {
		result.Token, err = uc.JwtConfig.GenerateToken(result.Id, true)
	} else {
		result.Token, err = uc.JwtConfig.GenerateToken(result.Id, false)
	}

	if err != nil {
		return Domain{}, err
	}
	return result, nil
}
