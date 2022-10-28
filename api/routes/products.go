package routes

import (
	"github.com/SilviaPabon/buenavida-backend/controllers"
	"github.com/labstack/echo/v4"
)

// SetupProductsRoutes initialize user routes
func SetupProductsRoutes(e *echo.Echo) {
	// Get all products
	e.GET("/api/products", controllers.HandleProductsGet)
	// Get products by page
	e.GET("/api/products/:page", controllers.HandleProductsPagination)
	// Search product fron text
	e.POST("/api/products/search", controllers.HandleProductsSearch)

	//Get details product
	e.GET("/api/product/:id", controllers.GetFromID)
}
