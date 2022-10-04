package routes

import "github.com/labstack/echo/v4"

// Creating a group of routes for the products.
func (r *Routes) ProductRoutes(e *echo.Echo) {
	products := e.Group("/products")
	products.POST("", r.controller.AddProduct)
}
