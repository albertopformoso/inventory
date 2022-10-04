package controller

import (
	"github.com/albertopformoso/inventory/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Controller interface {
	RegisterUser(c echo.Context) error
	LoginUser(c echo.Context) error
}

type controller struct {
	service       service.Service
	dataValidator *validator.Validate
}

func New(service service.Service) Controller {
	return &controller{
		service:       service,
		dataValidator: validator.New(),
	}
}
