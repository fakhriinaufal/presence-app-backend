package schedules

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"presence-app-backend/business/schedules"
	"presence-app-backend/controllers"
	"presence-app-backend/controllers/schedules/requests"
	"presence-app-backend/controllers/schedules/responses"
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
