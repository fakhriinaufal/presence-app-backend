package presences

import (
	"context"
	"presence-app-backend/business/schedules"
	"presence-app-backend/business/users"
	"strconv"
	"strings"
	"time"
)

type presenceUsecase struct {
	repo           Repository
	userRepo       users.Repository
	scheduleRepo   schedules.Repository
	contextTimeout time.Duration
}

func NewPresenceUsecase(pr Repository, ur users.Repository, sr schedules.Repository, timeout time.Duration) Usecase {
	return &presenceUsecase{
		pr,
		ur,
		sr,
		timeout,
	}
}

func (pu presenceUsecase) Store(ctx context.Context, domain *Domain) (Domain, error) {
	// check is user exist
	_, err := pu.userRepo.GetById(ctx, domain.UserId)
	if err != nil {
		return Domain{}, err
	}

	schedule, err := pu.scheduleRepo.GetById(ctx, domain.ScheduleId)
	if err != nil {
		return Domain{}, err
	}

	// check is user late or not
	currentTime := time.Now()
	var presenceSchedule time.Time
	LATE_TIME_MINUTE := 15
	var status string
	switch domain.Type {
	case "in":
		{
			scheduleTime := strings.Split(schedule.InTime, ":")
			presenceHour, _ := strconv.Atoi(scheduleTime[0])
			presenceMinute, _ := strconv.Atoi(scheduleTime[1])
			presenceSecond, _ := strconv.Atoi(scheduleTime[2])

			presenceSchedule = time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), presenceHour, presenceMinute, presenceSecond, 0, time.UTC)
			presenceSchedule.Add(time.Minute * time.Duration(LATE_TIME_MINUTE))

			isLate := currentTime.After(presenceSchedule)
			if isLate {
				status = "late"
			} else {
				status = "on time"
			}
		}
	case "out":
		{
			scheduleTime := strings.Split(schedule.OutTime, ":")
			presenceHour, _ := strconv.Atoi(scheduleTime[0])
			presenceMinute, _ := strconv.Atoi(scheduleTime[1])
			presenceSecond, _ := strconv.Atoi(scheduleTime[2])

			presenceSchedule = time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), presenceHour, presenceMinute, presenceSecond, 0, time.UTC)
			presenceSchedule.Add(time.Minute * time.Duration(LATE_TIME_MINUTE))

			isLate := currentTime.After(presenceSchedule)
			if isLate {
				status = "late"
			} else {
				status = "on time"
			}
		}
	}
	domain.Status = status
	result, err := pu.repo.Store(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return result, nil
}

func (pu presenceUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	result, err := pu.repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return result, nil
}

func (pu presenceUsecase) GetById(ctx context.Context, id int) (Domain, error) {
	result, err := pu.repo.GetById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	return result, nil
}

func (pu presenceUsecase) Update(ctx context.Context, domain *Domain) (Domain, error) {
	existedPresence, err := pu.repo.GetById(ctx, domain.Id)
	if err != nil {
		return Domain{}, err
	}
	domain.CreatedAt = existedPresence.UpdatedAt
	result, err := pu.repo.Update(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return result, nil
}
