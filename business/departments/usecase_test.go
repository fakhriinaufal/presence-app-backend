package departments_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"presence-app-backend/business/departments"
	mockDepartmentRepository "presence-app-backend/business/departments/mocks"
	"testing"
	"time"
)

var (
	departmentRepository mockDepartmentRepository.Repository
	departmentService departments.Usecase
	departmentDomain departments.Domain
)

func setup() {
	departmentService = departments.NewDepartmentUsecase(&departmentRepository, time.Hour*1)
	departmentDomain = departments.Domain{
		ID:          0,
		Name:        "Marketing",
		Description: "Share poster",
	}
}

func TestGetAll(t *testing.T) {
	setup()
	departmentRepository.On("GetAll", mock.Anything).Return([]departments.Domain{departmentDomain}, nil).Once()
	departmentRepository.On("GetAll", mock.Anything).Return([]departments.Domain{}, errors.New("failed")).Once()

	t.Run("Test Case 1 | Valid Get All", func(t *testing.T) {
		result, err := departmentService.GetAll(context.Background())

		assert.Nil(t, err)
		assert.Equal(t, "Marketing", result[0].Name)
	})

	t.Run("Test Case 2 | Invalid Get All", func(t *testing.T) {
		result, err := departmentService.GetAll(context.Background())

		assert.Equal(t, []departments.Domain{}, result)
		assert.Equal(t, errors.New("failed"), err)
	})
}

func TestStore(t *testing.T){
	setup()
	departmentRepository.On("Store", mock.Anything, mock.AnythingOfType("*departments.Domain")).Return(departmentDomain, nil).Once()
	departmentRepository.On("Store", mock.Anything, mock.AnythingOfType("*departments.Domain")).Return(departments.Domain{}, errors.New("failed")).Once()

	t.Run("Test Case 1 | Valid Store", func(t *testing.T) {
		result, err := departmentService.Store(context.Background(), &departmentDomain)

		assert.Nil(t, err)
		assert.Equal(t, departmentDomain, result)
	})

	t.Run("Test Case 2 | Invalid Store", func(t *testing.T) {
		result, err := departmentService.Store(context.Background(), &departmentDomain)

		assert.Equal(t, errors.New("failed"), err)
		assert.Equal(t, departments.Domain{}, result)
	})
}

func TestGetById(t *testing.T) {
	setup()
	departmentRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(departmentDomain, nil).Once()
	departmentRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(departments.Domain{}, errors.New("id doesn't exist")).Once()

	t.Run("Test Case 1 | Valid GetById", func(t *testing.T) {
		result, err := departmentService.GetById(context.Background(), 0)

		assert.Nil(t, err)
		assert.Equal(t, departmentDomain, result)
	})

	t.Run("Test Case 2 | Invalid GetById, Id doesn't exist", func(t *testing.T) {
		result, err := departmentService.GetById(context.Background(), 0)

		assert.Equal(t, errors.New("id doesn't exist"), err)
		assert.Equal(t, departments.Domain{}, result)
	})
}

func TestUpdate(t *testing.T){
	setup()
	t.Run("Test Case 1 | Valid Update", func(t *testing.T) {
		departmentRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(departmentDomain, nil).Once()
		departmentRepository.On("Update", mock.Anything, mock.AnythingOfType("*departments.Domain")).Return(departmentDomain, nil).Once()
		result, err := departmentService.Update(context.Background(), &departmentDomain)

		assert.Nil(t, err)
		assert.Equal(t, departmentDomain, result)
	})

	t.Run("Test Case 2 | Invalid Update, Department Doesn't Exist", func(t *testing.T) {
		departmentRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(departments.Domain{}, errors.New("not found")).Once()
		result, err := departmentService.Update(context.Background(), &departmentDomain)

		assert.Equal(t, errors.New("not found"), err)
		assert.Equal(t, departments.Domain{}, result)
	})

	t.Run("Test Case 3 | Invalid Update, error on database", func(t *testing.T) {
		departmentRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(departmentDomain, nil).Once()
		departmentRepository.On("Update", mock.Anything, mock.AnythingOfType("*departments.Domain")).Return(departments.Domain{}, errors.New("database error")).Once()

		result, err := departmentService.Update(context.Background(), &departmentDomain)

		assert.Equal(t, errors.New("database error"), err)
		assert.Equal(t, departments.Domain{}, result)
	})
}

func TestDelete(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Success delete", func(t *testing.T) {
		departmentRepository.On("Delete", mock.Anything, mock.AnythingOfType("int")).Return(nil).Once()
		err := departmentService.Delete(context.Background(), 0)
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Failed delete, id doesn't exist", func(t *testing.T) {
		departmentRepository.On("Delete", mock.Anything, mock.AnythingOfType("int")).Return(errors.New("failed")).Once()
		err := departmentService.Delete(context.Background(), 0)
		assert.Equal(t, errors.New("failed"), err)

	})
}