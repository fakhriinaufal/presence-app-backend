package routes

import (
	"github.com/labstack/echo/v4"
	"presence-app-backend/controllers/departments"
	"presence-app-backend/controllers/presences"
	"presence-app-backend/controllers/schedules"
	"presence-app-backend/controllers/users"
)

type ControllerList struct {
	DepartmentController departments.DepartmentController
	UserController       users.UserController
	ScheduleController   schedules.ScheduleController
	PresenceController   presences.PresenceController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {

	departmentRoute := e.Group("departments")
	departmentRoute.GET("", cl.DepartmentController.GetAll)
	departmentRoute.POST("", cl.DepartmentController.Store)
	departmentRoute.GET("/:id", cl.DepartmentController.GetById)
	departmentRoute.PUT("/:id", cl.DepartmentController.Update)
	departmentRoute.DELETE("/:id", cl.DepartmentController.Delete)

	userRoute := e.Group("userRoute")
	userRoute.POST("", cl.UserController.Store)
	userRoute.GET("", cl.UserController.GetAll)
	userRoute.GET("/:id", cl.UserController.GetById)
	userRoute.PUT("/:id", cl.UserController.Update)
	userRoute.DELETE("/:id", cl.UserController.Delete)

	scheduleRoute := e.Group("schedules")
	scheduleRoute.POST("", cl.ScheduleController.Store)
	scheduleRoute.GET("", cl.ScheduleController.GetAll)
	scheduleRoute.GET("/:id", cl.ScheduleController.GetById)
	scheduleRoute.PUT("/:id", cl.ScheduleController.Update)
	scheduleRoute.DELETE("/:id", cl.ScheduleController.Delete)

	presenceRoute := e.Group("presences")
	presenceRoute.POST("", cl.PresenceController.Create)
}
