package users_test

import (
	"context"
	"errors"
	_middleware "presence-app-backend/app/middlewares"
	"presence-app-backend/business/departments"
	_mockDepartmentRepo "presence-app-backend/business/departments/mocks"
	"presence-app-backend/business/users"
	_mockUserRepo "presence-app-backend/business/users/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	departmentRepository _mockDepartmentRepo.Repository
	userRepository       _mockUserRepo.Repository
	userDomain           users.Domain
	departmentDomain     departments.Domain
	userService          users.Usecase
)

func setup() {
	configJWT := _middleware.ConfigJWT{
		SecretJWT:       "123456",
		ExpiresDuration: 2,
	}
	userService = users.NewUserUsecase(&userRepository, &departmentRepository, time.Hour*1, configJWT)
	userDomain = users.Domain{
		Id:           1,
		DepartmentId: 1,
		Name:         "John",
		Password:     "apple",
		Email:        "john@doe.com",
	}
	departmentDomain = departments.Domain{
		ID:   1,
		Name: "admin",
	}
}

func TestStore(t *testing.T) {
	setup()
	t.Run("Test 1 | Department Not found", func(t *testing.T) {
		departmentRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(departments.Domain{}, errors.New("department not found")).Once()

		result, err := userService.Store(context.Background(), &userDomain)

		assert.Equal(t, users.Domain{}, result)
		assert.Equal(t, errors.New("department not found"), err)
	})
}

func TestGetAll(t *testing.T) {
	setup()
	t.Run("Test 1 | Success Get All Users", func(t *testing.T) {
		userRepository.On("GetAll").Return([]users.Domain{userDomain}, nil).Once()

		result, err := userService.GetAll()

		assert.Equal(t, []users.Domain{userDomain}, result)
		assert.Nil(t, err)
	})

	t.Run("Test 2 | Failed Get All Users", func(t *testing.T) {
		userRepository.On("GetAll").Return([]users.Domain{}, errors.New("failed")).Once()

		result, err := userService.GetAll()

		assert.Equal(t, []users.Domain{}, result)
		assert.Equal(t, errors.New("failed"), err)
	})
}

func TestGetById(t *testing.T) {
	setup()
	t.Run("Test 1 | Success get user", func(t *testing.T) {
		userRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(userDomain, nil).Once()

		result, err := userService.GetById(context.Background(), 1)

		assert.Equal(t, userDomain, result)
		assert.Nil(t, err)
	})

	t.Run("Test 2 | User doesn't exist", func(t *testing.T) {
		userRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(users.Domain{}, errors.New("user doesn't exist")).Once()

		result, err := userService.GetById(context.Background(), 1)

		assert.Equal(t, users.Domain{}, result)
		assert.Equal(t, errors.New("user doesn't exist"), err)
	})
}

func TestDelete(t *testing.T) {
	setup()
	t.Run("Test 1 | Success delete user", func(t *testing.T) {
		userRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(userDomain, nil).Once()
		userRepository.On("Delete", mock.Anything, mock.AnythingOfType("int")).Return(nil).Once()

		err := userService.Delete(context.Background(), 1)

		assert.Nil(t, err)
	})

	t.Run("Test 2 | User doesn't exist", func(t *testing.T) {
		userRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(users.Domain{}, errors.New("user doesn't exist")).Once()

		err := userService.Delete(context.Background(), 1)

		assert.Equal(t, errors.New("user doesn't exist"), err)
	})

	t.Run("Test 3 | Failed delete user", func(t *testing.T) {
		userRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(userDomain, nil).Once()
		userRepository.On("Delete", mock.Anything, mock.AnythingOfType("int")).Return(errors.New("failed")).Once()

		err := userService.Delete(context.Background(), 1)

		assert.Equal(t, errors.New("failed"), err)
	})
}

func TestUpdate(t *testing.T) {
	setup()
	t.Run("Test 1 | Success Update", func(t *testing.T) {
		userRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(userDomain, nil).Once()
		userRepository.On("Update", mock.Anything, mock.AnythingOfType("*users.Domain")).Return(userDomain, nil).Once()

		result, err := userService.Update(context.Background(), &userDomain, 1)

		assert.Equal(t, userDomain, result)
		assert.Nil(t, err)
	})
}
