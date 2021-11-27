package presences

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"presence-app-backend/business/presences"
	"presence-app-backend/controllers"
	"presence-app-backend/controllers/presences/requests"
	"presence-app-backend/controllers/presences/responses"
	"strconv"
)

type PresenceController struct {
	presenceUsecase presences.Usecase
}

func NewPresenceController(presenceUC presences.Usecase) *PresenceController {
	return &PresenceController{
		presenceUsecase: presenceUC,
	}
}

func (ctrl *PresenceController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := requests.Presence{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	reqDomain := req.ToDomain()
	response, err := ctrl.presenceUsecase.Store(ctx, reqDomain)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, map[string]interface{}{
		"presence": responses.FromDomain(response),
	})
}

func (ctrl *PresenceController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	result, err := ctrl.presenceUsecase.GetAll(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	var response []responses.Presence

	for _, val := range result {
		response = append(response, responses.FromDomain(val))
	}

	return controllers.NewSuccessResponse(c, map[string]interface{}{
		"presences": response,
	})
}

func (ctrl *PresenceController) GetById(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.presenceUsecase.GetById(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, map[string]interface{}{
		"presence": responses.FromDomain(result),
	})
}
