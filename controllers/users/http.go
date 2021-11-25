package users

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"presence-app-backend/business/users"
	"presence-app-backend/controllers"
	"presence-app-backend/controllers/users/request"
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

	user, err = controller.UserUsecase.Store(&user)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessCreatedResponse(c, map[string]interface{}{
		"user": user,
	})


}