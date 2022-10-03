package routes

import (
	"github.com/albertopformoso/inventory/internal/controller"
	"github.com/labstack/echo/v4"
)

type Routes struct {
	controller controller.Controller
}

func New(controller controller.Controller) *Routes {
	return &Routes{controller: controller}
}

func (r *Routes) Start(e *echo.Echo, address string) error {
	r.UserRoutes(e)
	return e.Start(address)
}
