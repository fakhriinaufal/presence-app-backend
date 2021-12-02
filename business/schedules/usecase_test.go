package schedules_test

import (
	"context"
	"errors"
	"presence-app-backend/business/departments"
	mockDepartmentRepository "presence-app-backend/business/departments/mocks"
	"presence-app-backend/business/schedules"
	mockScheduleRepository "presence-app-backend/business/schedules/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	scheduleRepository   mockScheduleRepository.Repository
	departmentRepository mockDepartmentRepository.Repository
	scheduleService      schedules.Usecase
	scheduleDomain       schedules.Domain
	departmentDomain     departments.Domain
)

func setup() {
	scheduleService = schedules.NewScheduleUsecase(&scheduleRepository, &departmentRepository, time.Hour*1)
	scheduleDomain = schedules.Domain{
		Id:           0,
		DepartmentId: 0,
		InTime:       "07:00:00",
		OutTime:      "15:00:00",
	}
	departmentDomain = departments.Domain{
		ID:          0,
		Name:        "Marketing",
		Description: "Share poster",
	}
}

func TestGetAll(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid Get All", func(t *testing.T) {
		scheduleRepository.On("GetAll", mock.Anything).Return([]schedules.Domain{scheduleDomain}, nil).Once()

		result, err := scheduleService.GetAll(context.Background())

		assert.Nil(t, err)
		assert.Equal(t, []schedules.Domain{scheduleDomain}, result)
	})

	t.Run("Test Case 2 | Invalid Get All", func(t *testing.T) {
		scheduleRepository.On("GetAll", mock.Anything).Return([]schedules.Domain{}, errors.New("failed")).Once()

		result, err := scheduleService.GetAll(context.Background())

		assert.Equal(t, errors.New("failed"), err)
		assert.Equal(t, []schedules.Domain{}, result)
	})
}

func TestStore(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Success Store", func(t *testing.T) {
		departmentRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(departmentDomain, nil).Once()
		scheduleRepository.On("Store", mock.Anything, mock.AnythingOfType("*schedules.Domain")).Return(scheduleDomain, nil).Once()

		result, err := scheduleService.Store(context.Background(), &scheduleDomain)

		assert.Nil(t, err)
		assert.Equal(t, scheduleDomain, result)
	})

	t.Run("Test Case 2 | Department Not Found", func(t *testing.T) {
		departmentRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(departments.Domain{}, errors.New("department doesn't exist")).Once()

		result, err := scheduleService.Store(context.Background(), &scheduleDomain)

		assert.Equal(t, errors.New("department doesn't exist"), err)
		assert.Equal(t, schedules.Domain{}, result)
	})

	t.Run("Test Case 3 | Failed Store", func(t *testing.T) {
		departmentRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(departmentDomain, nil).Once()
		scheduleRepository.On("Store", mock.Anything, mock.AnythingOfType("*schedules.Domain")).Return(schedules.Domain{}, errors.New("failed")).Once()

		result, err := scheduleService.Store(context.Background(), &scheduleDomain)

		assert.Equal(t, errors.New("failed"), err)
		assert.Equal(t, schedules.Domain{}, result)
	})
}

func TestGetById(t *testing.T) {
	setup()

	t.Run("Test Case 1 | Schedule Found", func(t *testing.T) {
		scheduleRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(scheduleDomain, nil).Once()

		result, err := scheduleService.GetById(context.Background(), 0)

		assert.Nil(t, err)
		assert.Equal(t, scheduleDomain, result)
	})

	t.Run("Test Case 2 | Schedule Not Found", func(t *testing.T) {
		scheduleRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(schedules.Domain{}, errors.New("failed")).Once()

		result, err := scheduleService.GetById(context.Background(), 0)

		assert.Equal(t, errors.New("failed"), err)
		assert.Equal(t, schedules.Domain{}, result)
	})
}

func TestUpdate(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Success Update", func(t *testing.T) {
		scheduleRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(scheduleDomain, nil).Once()
		scheduleRepository.On("Update", mock.Anything, mock.AnythingOfType("*schedules.Domain")).Return(scheduleDomain, nil).Once()

		result, err := scheduleService.Update(context.Background(), &scheduleDomain)

		assert.Nil(t, err)
		assert.Equal(t, scheduleDomain, result)
	})

	t.Run("Test Case 2 | Schedule Not Found", func(t *testing.T) {
		scheduleRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(schedules.Domain{}, errors.New("schedule doesn't exist")).Once()

		result, err := scheduleService.Update(context.Background(), &scheduleDomain)

		assert.Equal(t, errors.New("schedule doesn't exist"), err)
		assert.Equal(t, schedules.Domain{}, result)
	})

	t.Run("Test Case 3 | Failed Update", func(t *testing.T) {
		scheduleRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(scheduleDomain, nil).Once()
		scheduleRepository.On("Update", mock.Anything, mock.AnythingOfType("*schedules.Domain")).Return(schedules.Domain{}, errors.New("failed")).Once()

		result, err := scheduleService.Update(context.Background(), &scheduleDomain)

		assert.Equal(t, errors.New("failed"), err)
		assert.Equal(t, schedules.Domain{}, result)
	})
}

func TestDelete(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Success Delete", func(t *testing.T) {
		scheduleRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(scheduleDomain, nil).Once()
		scheduleRepository.On("Delete", mock.Anything, mock.AnythingOfType("int")).Return(nil).Once()

		err := scheduleService.Delete(context.Background(), 0)

		assert.Nil(t, err)

	})

	t.Run("Test Case 1 | Schedule Doesn't Exist", func(t *testing.T) {
		scheduleRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(schedules.Domain{}, errors.New("schedule doesn't exist")).Once()

		err := scheduleService.Delete(context.Background(), 0)

		assert.Equal(t, errors.New("schedule doesn't exist"), err)
	})
}
