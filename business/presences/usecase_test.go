package presences_test

import (
	"context"
	"errors"
	"presence-app-backend/business/presences"
	mockPresenceRepository "presence-app-backend/business/presences/mocks"
	"presence-app-backend/business/schedules"
	mockScheduleRepository "presence-app-backend/business/schedules/mocks"
	"presence-app-backend/business/users"
	mockUserRepository "presence-app-backend/business/users/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	presenceRepository mockPresenceRepository.Repository
	userRepository     mockUserRepository.Repository
	scheduleRepository mockScheduleRepository.Repository
	presenceDomain     presences.Domain
	presenceService    presences.Usecase
	userDomain         users.Domain
	scheduleDomain     schedules.Domain
)

func setup() {
	presenceService = presences.NewPresenceUsecase(&presenceRepository, &userRepository, &scheduleRepository, time.Hour*1)
	presenceDomain = presences.Domain{
		Id:         1,
		UserId:     1,
		ScheduleId: 1,
		Type:       "in",
		Status:     "late",
	}
	userDomain = users.Domain{
		Id:           1,
		DepartmentId: 1,
		Name:         "John",
		Email:        "john@doe.com",
		Dob:          "2007-11-01",
	}
	scheduleDomain = schedules.Domain{
		Id:           1,
		DepartmentId: 1,
		InTime:       "07:00:00",
		OutTime:      "15:00:00",
	}

}

func TestStore(t *testing.T) {
	presencePayload := presences.Domain{
		UserId:     1,
		ScheduleId: 1,
		Type:       "in",
	}
	setup()
	t.Run("Test 1 | User doesn't exist", func(t *testing.T) {
		userRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(users.Domain{}, errors.New("user doesn't exist")).Once()

		result, err := presenceService.Store(context.Background(), &presencePayload)

		assert.Equal(t, presences.Domain{}, result)
		assert.Equal(t, errors.New("user doesn't exist"), err)
	})

	t.Run("Test 2 | Schedule database error", func(t *testing.T) {
		userRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(userDomain, nil).Once()
		scheduleRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(schedules.Domain{}, errors.New("failed")).Once()

		result, err := presenceService.Store(context.Background(), &presencePayload)

		assert.Equal(t, presences.Domain{}, result)
		assert.Equal(t, errors.New("failed"), err)
	})

	t.Run("Test 3 | Succes store", func(t *testing.T) {
		userRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(userDomain, nil).Once()
		scheduleRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(scheduleDomain, nil).Once()
		presenceRepository.On("Store", mock.Anything, mock.AnythingOfType("*presences.Domain")).Return(presenceDomain, nil).Once()

		result, err := presenceService.Store(context.Background(), &presencePayload)

		assert.Equal(t, presenceDomain, result)
		assert.Nil(t, err)
	})
}

func TestGetAll(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Success get all", func(t *testing.T) {
		presenceRepository.On("GetAll", mock.Anything).Return([]presences.Domain{presenceDomain}, nil).Once()

		result, err := presenceService.GetAll(context.Background())

		assert.Nil(t, err)
		assert.Equal(t, []presences.Domain{presenceDomain}, result)
	})

	t.Run("Test Case 2 | Failed get all", func(t *testing.T) {
		presenceRepository.On("GetAll", mock.Anything).Return([]presences.Domain{}, errors.New("failed")).Once()

		result, err := presenceService.GetAll(context.Background())

		assert.Equal(t, []presences.Domain{}, result)
		assert.Equal(t, errors.New("failed"), err)
	})
}

func TestGetById(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Success get presence", func(t *testing.T) {
		presenceRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(presenceDomain, nil).Once()

		result, err := presenceService.GetById(context.Background(), 1)

		assert.Nil(t, err)
		assert.Equal(t, presenceDomain, result)
	})

	t.Run("Test Case 2 | Presence doesn't exist", func(t *testing.T) {
		presenceRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(presences.Domain{}, errors.New("presence doesn't exist")).Once()

		result, err := presenceService.GetById(context.Background(), 1)

		assert.Equal(t, presences.Domain{}, result)
		assert.Equal(t, errors.New("presence doesn't exist"), err)
	})
}

func TestUpdate(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Successful update", func(t *testing.T) {
		presenceRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(presenceDomain, nil).Once()
		presenceRepository.On("Update", mock.Anything, mock.AnythingOfType("*presences.Domain")).Return(presenceDomain, nil).Once()

		result, err := presenceService.Update(context.Background(), &presenceDomain)

		assert.Nil(t, err)
		assert.Equal(t, presenceDomain, result)
	})

	t.Run("Test Case 2 | Presence Doesn't Exist", func(t *testing.T) {
		presenceRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(presences.Domain{}, errors.New("presence doesn't exist")).Once()

		result, err := presenceService.Update(context.Background(), &presenceDomain)

		assert.Equal(t, presences.Domain{}, result)
		assert.Equal(t, errors.New("presence doesn't exist"), err)
	})

	t.Run("Test Case 3 | Failed", func(t *testing.T) {
		presenceRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(presenceDomain, nil).Once()
		presenceRepository.On("Update", mock.Anything, mock.AnythingOfType("*presences.Domain")).Return(presences.Domain{}, errors.New("failed")).Once()

		result, err := presenceService.Update(context.Background(), &presenceDomain)

		assert.Equal(t, errors.New("failed"), err)
		assert.Equal(t, presences.Domain{}, result)
	})
}

func TestDelete(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Success Delete", func(t *testing.T) {
		presenceRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(presenceDomain, nil).Once()
		presenceRepository.On("Delete", mock.Anything, mock.AnythingOfType("int")).Return(nil).Once()

		err := presenceService.Delete(context.Background(), 1)

		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Presence doesn't exist", func(t *testing.T) {
		presenceRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(presences.Domain{}, errors.New("presence doesn't exist")).Once()

		err := presenceService.Delete(context.Background(), 1)

		assert.Equal(t, errors.New("presence doesn't exist"), err)
	})

	t.Run("Test Case 3 | Failed", func(t *testing.T) {
		presenceRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(presenceDomain, nil).Once()
		presenceRepository.On("Delete", mock.Anything, mock.AnythingOfType("int")).Return(errors.New("failed")).Once()

		err := presenceService.Delete(context.Background(), 1)

		assert.Equal(t, errors.New("failed"), err)
	})
}
