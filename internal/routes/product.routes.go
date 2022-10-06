package routes

import "github.com/labstack/echo/v4"

// Creating a group of routes for the products.
func (r *routes) productRoutes(e *echo.Echo) {
	products := e.Group("/products")
	products.POST("", r.controller.AddProduct)
	products.GET("", r.controller.GetProducts)
	products.GET("/:id", r.controller.GetProduct)
}
