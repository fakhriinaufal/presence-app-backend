package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"presence-app-backend/app/middlewares"
	"presence-app-backend/controllers/departments"
	"presence-app-backend/controllers/presences"
	"presence-app-backend/controllers/schedules"
	"presence-app-backend/controllers/users"
)

type ControllerList struct {
	JwtConfig            middleware.JWTConfig
	DepartmentController departments.DepartmentController
	UserController       users.UserController
	ScheduleController   schedules.ScheduleController
	PresenceController   presences.PresenceController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {

	departmentRoute := e.Group("departments")
	departmentRoute.Use(middleware.JWTWithConfig(cl.JwtConfig))
	departmentRoute.Use(middlewares.IsAdmin())
	departmentRoute.GET("", cl.DepartmentController.GetAll)
	departmentRoute.POST("", cl.DepartmentController.Store)
	departmentRoute.GET("/:id", cl.DepartmentController.GetById)
	departmentRoute.PUT("/:id", cl.DepartmentController.Update)
	departmentRoute.DELETE("/:id", cl.DepartmentController.Delete)

	e.POST("users/login", cl.UserController.Login)
	userRoute := e.Group("users")
	userRoute.POST("", cl.UserController.Store)
	userRoute.GET("", cl.UserController.GetAll)
	userRoute.GET("/:id", cl.UserController.GetById)
	userRoute.PUT("/:id", cl.UserController.Update)
	userRoute.DELETE("/:id", cl.UserController.Delete)

	scheduleRoute := e.Group("schedules")
	scheduleRoute.Use(middleware.JWTWithConfig(cl.JwtConfig))
	scheduleRoute.Use(middlewares.IsAdmin())
	scheduleRoute.POST("", cl.ScheduleController.Store)
	scheduleRoute.GET("", cl.ScheduleController.GetAll)
	scheduleRoute.GET("/:id", cl.ScheduleController.GetById)
	scheduleRoute.PUT("/:id", cl.ScheduleController.Update)
	scheduleRoute.DELETE("/:id", cl.ScheduleController.Delete)

	presenceRoute := e.Group("presences")
	presenceRoute.Use(middleware.JWTWithConfig(cl.JwtConfig))
	presenceRoute.POST("", cl.PresenceController.Store)
	presenceRoute.GET("", cl.PresenceController.GetAll)
	presenceRoute.GET("/:id", cl.PresenceController.GetById)
	presenceRoute.PUT("/:id", cl.PresenceController.Update)
	presenceRoute.DELETE("/:id", cl.PresenceController.Delete)
}
