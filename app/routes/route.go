package routes

import (
	"github.com/labstack/echo/v4"
	"presence-app-backend/controllers/departments"
)

type ControllerList struct {
	DepartmentController departments.DepartmentController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.GET("departments", cl.DepartmentController.GetAll)
}