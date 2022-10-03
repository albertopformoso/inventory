package api

import (
	"github.com/albertopformoso/inventory/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type API struct {
	service service.Service
	dataValidator *validator.Validate
}

func New(service service.Service) *API {
	return &API{
		service: service,
		dataValidator: validator.New(),
	}
}

func (api *API) Start(e *echo.Echo, address string) error {
	api.RegisterRoutes(e)
	return e.Start(address)
}
