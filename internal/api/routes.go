package api

import "github.com/labstack/echo/v4"

func (api *API) RegisterRoutes(e *echo.Echo) {
	users := e.Group("/users")
	users.POST("/register", api.RegisterUser)

	products := e.Group("/products")
	products.POST("/register", nil)
}