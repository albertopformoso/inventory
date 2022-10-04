package routes

import (
	"github.com/albertopformoso/inventory/internal/controller"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

type Routes struct {
	controller controller.Controller
}

func New(controller controller.Controller) *Routes {
	return &Routes{controller: controller}
}

// Initialize the routes.
func (r *Routes) RegisterRoutes(e *echo.Echo) {
	r.UserRoutes(e)
	r.ProductRoutes(e)
}

// Setting up the middlewares for the server.
func (r *Routes) ConfigMiddlewares(e *echo.Echo) {
	e.Use(
		middleware.Recover(),
		middleware.Logger(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{"http://127.0.0.1:8080"},
			AllowMethods:     []string{echo.POST},
			AllowHeaders:     []string{echo.HeaderContentType},
			AllowCredentials: true,
		}))
}

// Starting the server.
func (r *Routes) Start(e *echo.Echo, address string) error {
	r.ConfigMiddlewares(e)
	r.RegisterRoutes(e)

	return e.Start(address)
}
