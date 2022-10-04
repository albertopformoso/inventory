package controller

import (
	"github.com/albertopformoso/inventory/internal/service"
	"github.com/albertopformoso/inventory/logger"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type Controller interface {
	RegisterUser(c echo.Context) error
	LoginUser(c echo.Context) error

	AddProduct(c echo.Context) error
}

type controller struct {
	service       service.Service
	log           *zerolog.Logger
	dataValidator *validator.Validate
}

func New(service service.Service) Controller {
	return &controller{
		service:       service,
		log:           logger.New(),
		dataValidator: validator.New(),
	}
}
