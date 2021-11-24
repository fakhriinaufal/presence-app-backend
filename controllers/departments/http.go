package departments

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	dpt "presence-app-backend/business/departments"
	"presence-app-backend/controllers"
	"presence-app-backend/controllers/departments/requests"
	"presence-app-backend/controllers/departments/responses"
)

type DepartmentController struct {
	DepartmentUsecase dpt.Usecase
}

func NewDepartmentController(departmentUsecase dpt.Usecase) *DepartmentController {
	return &DepartmentController{
		departmentUsecase,
	}
}

func (d DepartmentController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	departments, err := d.DepartmentUsecase.GetAll(ctx)

	var returnedDepartment []responses.DepartmentResponse

	for _, val := range departments {
		returnedDepartment = append(returnedDepartment, responses.FromDomain(val))
	}

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusOK, err)
	}
	if returnedDepartment == nil {
		return controllers.NewSuccessResponse(c, []int{})
	}
	return controllers.NewSuccessResponse(c, returnedDepartment)
}

func isDepartmentStoreValid(request *requests.DepartmentStore) (bool, error) {
	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		return false, err
	}
	return true, nil

}

func (d DepartmentController) Store(c echo.Context) error {
	var department requests.DepartmentStore
	var err error
	err = c.Bind(&department)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}


	isValid, err := isDepartmentStoreValid(&department)

	if !isValid {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	dept := department.ToDomain()
	dept, err = d.DepartmentUsecase.Store(ctx, &dept)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(
		c,
		map[string]interface{}{
			"department": dept,
		})
}
