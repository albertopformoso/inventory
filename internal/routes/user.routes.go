package routes

import "github.com/labstack/echo/v4"

// Creating a group of routes for the users.
func (r *routes) userRoutes(e *echo.Echo) {
	users := e.Group("/users")
	users.POST("/register", r.controller.RegisterUser)
	users.POST("/login", r.controller.LoginUser)
	users.POST("/role", r.controller.AddUserRole)
	users.DELETE("/role", r.controller.RemoveUserRole)
}
