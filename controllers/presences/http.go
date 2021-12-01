package presences

import (
	"github.com/go-playground/validator/v10"
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

func validatePayload(payload interface{}) error {
	validate := validator.New()
	return validate.Struct(payload)
}

func (ctrl *PresenceController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := requests.Presence{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	if err := validatePayload(req); err != nil {
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

func (ctrl *PresenceController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))

	var presenceRequest requests.PresenceUpdate
	if err := c.Bind(&presenceRequest); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	if err := validatePayload(presenceRequest); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	presenceDomain := presenceRequest.ToDomain()
	presenceDomain.Id = id

	result, err := ctrl.presenceUsecase.Update(ctx, presenceDomain)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, map[string]interface{}{
		"presence": responses.FromDomain(result),
	})
}

func (ctrl *PresenceController) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))

	if err := ctrl.presenceUsecase.Delete(ctx, id); err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, []int{})
}