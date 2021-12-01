package departments

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	dpt "presence-app-backend/business/departments"
	"presence-app-backend/controllers"
	"presence-app-backend/controllers/departments/requests"
	"presence-app-backend/controllers/departments/responses"
	"strconv"
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
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
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
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	isValid, err := isDepartmentStoreValid(&department)

	if !isValid {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()

	dept, err := d.DepartmentUsecase.Store(ctx, department.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(
		c,
		map[string]interface{}{
			"department": responses.FromDomain(dept),
		})
}

func (d DepartmentController) GetById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	ctx := c.Request().Context()
	result, err := d.DepartmentUsecase.GetById(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusNotFound, err)
	}
	return controllers.NewSuccessResponse(c, map[string]interface{}{
		"department": responses.FromDomain(result),
	})
}

func (d DepartmentController) Update(c echo.Context) error {
	var department requests.DepartmentStore
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	c.Bind(&department)

	isValid, err := isDepartmentStoreValid(&department)

	if !isValid {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	deptDomain := department.ToDomain()
	deptDomain.ID = id
	dept, err := d.DepartmentUsecase.Update(ctx, deptDomain)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusNotFound, err)
	}

	return controllers.NewSuccessResponse(c, map[string]interface{}{
		"department": responses.FromDomain(dept),
	})
}

func (d DepartmentController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	ctx := c.Request().Context()
	err := d.DepartmentUsecase.Delete(ctx, id)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusNotFound, err)
	}

	return controllers.NewSuccessResponse(c, map[string]interface{}{})
}
