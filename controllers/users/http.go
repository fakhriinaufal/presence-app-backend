package users

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"presence-app-backend/business/users"
	"presence-app-backend/controllers"
	"presence-app-backend/controllers/users/request"
	"presence-app-backend/controllers/users/responses"
	"strconv"
)

type UserController struct {
	UserUsecase users.Usecase
}

func NewUserController(usecase users.Usecase) *UserController {
	return &UserController{
		UserUsecase: usecase,
	}
}

func (controller UserController) Store(c echo.Context) error {
	//validate := validator.New()
	var userPayload request.UserPayload
	err := c.Bind(&userPayload)

	user := userPayload.ToDomain()

	ctx := c.Request().Context()

	user, err = controller.UserUsecase.Store(ctx, &user)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessCreatedResponse(c, map[string]interface{}{
		"user": user,
	})
}

func (controller UserController) GetAll(c echo.Context) error {
	usersFromUseCase, err := controller.UserUsecase.GetAll()

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	if usersFromUseCase == nil {
		return controllers.NewSuccessResponse(c, map[string]interface{}{
			"users": []int{},
		})
	}

	return controllers.NewSuccessResponse(c, map[string]interface{}{
		"users": responses.ToResponseList(&usersFromUseCase),
	})
}


func (controller UserController) GetById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()
	userFromUsecase, err := controller.UserUsecase.GetById(ctx, id)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, map[string]interface{}{
		"user": userFromUsecase,
	})
}

func (controller UserController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()
	var payload request.UserPayload
	err := c.Bind(&payload)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := payload.ToDomain()
	result, err := controller.UserUsecase.Update(ctx, &domainReq, id)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(result))
}
func (controller UserController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	err := controller.UserUsecase.Delete(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, []int{})
}
