package controller

import (
	"log"
	"net/http"

	"github.com/albertopformoso/inventory/encryption"
	"github.com/albertopformoso/inventory/internal/controller/dto"
	"github.com/albertopformoso/inventory/internal/model"
	"github.com/albertopformoso/inventory/internal/service"
	"github.com/labstack/echo/v4"
)

func (ctrl *controller) AddProduct(c echo.Context) error {
	// Get auth token from cookie
	cookie, err := c.Cookie("Authorization")
	if err != nil {
		log.Println(err)
		return c.JSONPretty(http.StatusUnauthorized, responseMsg{Message: "Unauthorized"}, "  ")
	}

	// Parse the jwt token
	claims, err := encryption.ParseLoginJWT(cookie.Value)
	if err != nil {
		log.Println(err)
		if err.Error() == "Token is expired" {
			return c.JSONPretty(http.StatusUnauthorized, responseMsg{Message: err.Error()}, "  ")
		}
		return c.JSONPretty(http.StatusUnauthorized, responseMsg{Message: "Unauthorized"}, "  ")
	}

	email := claims["email"].(string)

	// Get the payload from the request
	ctx := c.Request().Context()
	params := dto.AddProduct{}

	if err := c.Bind(&params); err != nil {
		log.Println(err)
		return c.JSONPretty(http.StatusBadRequest, responseMsg{Message: ErrInvalidRequest.Error()}, "  ")
	}

	if err := ctrl.dataValidator.Struct(params); err != nil {
		log.Println(err)
		return c.JSONPretty(http.StatusBadRequest, responseMsg{Message: ErrInvalidRequest.Error()}, "  ")
	}

	p := model.Product{
		Name:        params.Name,
		Description: params.Description,
		Price:       params.Price,
	}

	if err := ctrl.service.AddProduct(ctx, p, email); err != nil {
		log.Println(err)

		if err == service.ErrInvalidPermissions {
			return c.JSONPretty(http.StatusForbidden, responseMsg{Message: service.ErrInvalidPermissions.Error()}, "  ")
		}

		return c.JSONPretty(http.StatusInternalServerError, responseMsg{Message: ErrInternalServerError.Error()}, "  ")
	}

	return c.JSONPretty(http.StatusCreated, nil, "  ")
}
