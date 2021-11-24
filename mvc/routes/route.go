package routes

import (
	"net/http"
	c "presence-app-backend/mvc/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// hello world
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success",
			"data":    "Hello world",
		})
	})

	// deparments
	eDepartment := e.Group("/departments")
	eDepartment.GET("", c.GetDepartmentController)
	eDepartment.POST("", c.CreateDepartmentController)

	return e
}
