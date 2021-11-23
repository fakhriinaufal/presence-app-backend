package departments

import (
	"context"
	"time"
)

type DepartmentUsecase struct {
	Repo Repository
	contextTimeout time.Duration
}

func NewDepartmentUsecase(repo Repository, timeout time.Duration) Usecase {
	return &DepartmentUsecase{
		repo,
		timeout,
	}
}

func (uc *DepartmentUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	departments, err := uc.Repo.GetAll(ctx)

	if err != nil {
		return []Domain{}, err
	}

	return departments, nil
}