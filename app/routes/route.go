package routes

import (
	"github.com/labstack/echo/v4"
	"presence-app-backend/controllers/departments"
	"presence-app-backend/controllers/schedules"
	"presence-app-backend/controllers/users"
)

type ControllerList struct {
	DepartmentController departments.DepartmentController
	UserController users.UserController
	ScheduleController schedules.ScheduleController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.GET("departments", cl.DepartmentController.GetAll)
	e.POST("departments", cl.DepartmentController.Store)
	e.GET("departments/:id", cl.DepartmentController.GetById)
	e.PUT("departments/:id", cl.DepartmentController.Update)
	e.DELETE("departments/:id", cl.DepartmentController.Delete)

	users := e.Group("users")
	users.POST("", cl.UserController.Store)
	users.GET("", cl.UserController.GetAll)
	users.GET("/:id", cl.UserController.GetById)
	users.PUT("/:id", cl.UserController.Update)
	users.DELETE("/:id", cl.UserController.Delete)

	schedules := e.Group("schedules")
	schedules.POST("", cl.ScheduleController.Store)
	schedules.GET("", cl.ScheduleController.GetAll)
}
