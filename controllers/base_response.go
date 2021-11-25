package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type BaseResponse struct {
	Meta struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
	Data interface{} `json:"data"`
}

func NewSuccessResponse(c echo.Context, data interface{}) error {
	response := BaseResponse{}
	response.Meta.Status = http.StatusOK
	response.Meta.Message = "Success"
	response.Data = data

	return c.JSON(http.StatusOK, response)
}

func NewSuccessCreatedResponse(c echo.Context, data interface{}) error {
	response := BaseResponse{}
	response.Meta.Status = http.StatusCreated
	response.Meta.Message = "Success"
	response.Data = data

	return c.JSON(http.StatusOK, response)
}

func NewErrorResponse(c echo.Context, code int, err error) error {
	response := BaseResponse{}
	response.Meta.Status = code
	response.Meta.Message = err.Error()
	response.Data = nil
	return c.JSON(code, response)
}
