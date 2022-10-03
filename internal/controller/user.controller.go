package controller

import (
	"net/http"

	"github.com/albertopformoso/inventory/internal/controller/dto"
	"github.com/albertopformoso/inventory/internal/service"

	"github.com/labstack/echo/v4"
)

func (ctrl *controller) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	params := dto.RegisterUser{}

	if err := c.Bind(&params); err != nil {
		return c.JSONPretty(http.StatusBadRequest, responseMsg{Message: ErrInvalidRequest.Error()}, "  ")
	}

	if err := ctrl.dataValidator.Struct(params); err != nil {
		return c.JSONPretty(http.StatusBadRequest, responseMsg{Message: err.Error()}, "  ")
	}

	if err := ctrl.service.RegisterUser(ctx, params.Email, params.Name, params.Password); err != nil {
		if err == service.ErrUserAlreadyExists {
			return c.JSONPretty(http.StatusConflict, responseMsg{Message: err.Error()}, "  ")
		}

		return c.JSONPretty(http.StatusInternalServerError, responseMsg{Message: ErrInternalServerError.Error()}, "  ")
	}

	return c.JSONPretty(http.StatusCreated, nil, "  ")
}