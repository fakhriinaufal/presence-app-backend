package schedules

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"presence-app-backend/business/schedules"
	"presence-app-backend/controllers"
	"presence-app-backend/controllers/schedules/requests"
	"presence-app-backend/controllers/schedules/responses"
	"strconv"
)

type ScheduleController struct {
	scheduleUsecase schedules.Usecase
}

func NewScheduleController(scheduleUC schedules.Usecase) *ScheduleController {
	return &ScheduleController{
		scheduleUsecase: scheduleUC,
	}
}

func (ctrl *ScheduleController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := requests.Schedule{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	response, err := ctrl.scheduleUsecase.Store(ctx, req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessCreatedResponse(c, responses.FromDomain(response))
}

func (ctrl *ScheduleController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := ctrl.scheduleUsecase.GetAll(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	var result []responses.Schedule
	for _, val := range resp {
		result = append(result, responses.FromDomain(val))
	}

	return controllers.NewSuccessResponse(c, result)
}

func (ctrl *ScheduleController) GetById(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))

	resp, err := ctrl.scheduleUsecase.GetById(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusNotFound, err)
	}
	return controllers.NewSuccessResponse(c, map[string]interface{}{
		"schedule": responses.FromDomain(resp),
	})
}

func (ctrl *ScheduleController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	payload := requests.Schedule{}
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&payload); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	reqDomain := payload.ToDomain()
	reqDomain.Id = id
	result, err := ctrl.scheduleUsecase.Update(ctx, reqDomain)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, map[string]interface{}{
		"schedule": responses.FromDomain(result),
	})
}

func (ctrl *ScheduleController) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))

	err := ctrl.scheduleUsecase.Delete(ctx ,id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusNotFound, err)
	}
	return controllers.NewSuccessResponse(c, []int{})
}