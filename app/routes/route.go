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
	e.POST("departments", cl.DepartmentController.Store)
	e.GET("departments/:id", cl.DepartmentController.GetById)
	e.PUT("departments/:id", cl.DepartmentController.Update)
	e.DELETE("departments/:id", cl.DepartmentController.Delete)
}
