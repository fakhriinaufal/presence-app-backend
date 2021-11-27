package presences

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"presence-app-backend/business/presences"
	"presence-app-backend/controllers"
	"presence-app-backend/controllers/presences/requests"
	"presence-app-backend/controllers/presences/responses"
)

type PresenceController struct {
	presenceUsecase presences.Usecase
}

func NewPresenceController(presenceUC presences.Usecase) *PresenceController {
	return &PresenceController{
		presenceUsecase: presenceUC,
	}
}

func (ctrl *PresenceController) Create(c echo.Context) error {
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
