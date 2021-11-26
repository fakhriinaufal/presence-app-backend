package schedules

import (
	"context"
	"presence-app-backend/business/departments"
	"time"
)

type scheduleUsecase struct {
	scheduleRepository Repository
	departmentRepository departments.Repository
	contextTimeout time.Duration
}

func NewScheduleUsecase(sr Repository, dr departments.Repository, timeout time.Duration) Usecase {
	return &scheduleUsecase{
		sr,
		dr,
		timeout,
	}
}

func (su scheduleUsecase) Store(ctx context.Context, domain *Domain) (Domain, error) {
	_, err := su.departmentRepository.GetById(ctx, domain.DepartmentId)
	if err != nil {
		return Domain{}, err
	}

	result, err := su.scheduleRepository.Store(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return result, nil
}

func (su scheduleUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	schedules, err := su.scheduleRepository.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return schedules, nil
}

func (su scheduleUsecase) GetById(ctx context.Context, id int) (Domain, error) {
	result, err := su.scheduleRepository.GetById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	return result, nil
}