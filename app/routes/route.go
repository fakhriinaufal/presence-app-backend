package routes

import (
	"github.com/labstack/echo/v4"
	"presence-app-backend/controllers/departments"
	"presence-app-backend/controllers/users"
)

type ControllerList struct {
	DepartmentController departments.DepartmentController
	UserController users.UserController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.GET("departments", cl.DepartmentController.GetAll)
	e.POST("departments", cl.DepartmentController.Store)
	e.GET("departments/:id", cl.DepartmentController.GetById)
	e.PUT("departments/:id", cl.DepartmentController.Update)
	e.DELETE("departments/:id", cl.DepartmentController.Delete)

	users := e.Group("users")
	users.POST("", cl.UserController.Store)
}
