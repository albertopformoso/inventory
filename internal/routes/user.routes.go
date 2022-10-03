package routes

import "github.com/labstack/echo/v4"

func (r *Routes) UserRoutes(e *echo.Echo) {
	users := e.Group("/users")
	users.POST("/register", r.controller.RegisterUser)
}
