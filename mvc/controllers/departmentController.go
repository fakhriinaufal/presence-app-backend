package controllers

import (
	"net/http"
	"presence-app-backend/mvc/lib/database"
	"presence-app-backend/mvc/models"

	"github.com/labstack/echo/v4"
)

func GetDepartmentController(c echo.Context) error {
	departments, e := database.GetDepartments()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": e.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data": map[string]interface{}{
			"departments": departments,
		},
	})
}

func CreateDepartmentController(c echo.Context) error {
	var department models.Department
	c.Bind(&department)

	if e := database.CreateDepartment(&department); e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": e.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data": map[string]interface{}{
			"department": department,
		},
	})
}
