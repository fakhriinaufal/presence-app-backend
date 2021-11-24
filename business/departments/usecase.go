package departments

import (
	"context"
	"time"
)

type DepartmentUsecase struct {
	Repo           Repository
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

func (uc *DepartmentUsecase) Store(ctx context.Context, department *Domain) (Domain, error) {
	result, err := uc.Repo.Store(ctx, department)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (uc *DepartmentUsecase) GetById(ctx context.Context, id int) (Domain, error) {
	result, err := uc.Repo.GetById(ctx, id)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (uc *DepartmentUsecase) Update(ctx context.Context, department Domain, id int) (Domain, error) {
	result, err := uc.Repo.Update(ctx, department, id)

	if err != nil {
		return Domain{}, err
	}
	return result, nil
}

func (uc *DepartmentUsecase) Delete(ctx context.Context, id int) error {
	err := uc.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
