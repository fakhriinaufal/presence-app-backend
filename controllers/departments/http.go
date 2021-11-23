package departments

import (
	"github.com/labstack/echo/v4"
	"net/http"
	dpt "presence-app-backend/business/departments"
	"presence-app-backend/controllers"
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

func (departmentController DepartmentController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	departments, err := departmentController.DepartmentUsecase.GetAll(ctx)

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
