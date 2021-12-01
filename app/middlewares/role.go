package middlewares

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"presence-app-backend/controllers"
)

func IsAdmin() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := GetUser(c)
			if claims.IsAdmin {
				return hf(c)
			} else {
				return controllers.NewErrorResponse(c, http.StatusForbidden, errors.New("forbidden access"))
			}
		}
	}
}
