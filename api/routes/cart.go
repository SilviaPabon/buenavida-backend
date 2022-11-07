package routes

import (
	"github.com/SilviaPabon/buenavida-backend/controllers"
	"github.com/SilviaPabon/buenavida-backend/middlewares"
	"github.com/labstack/echo/v4"
)

// SetupCartRoutes create an manage cart routes
func SetupCartRoutes(e *echo.Echo) {
	// Add a new product to the cart
	e.POST("/api/cart", controllers.HandleCartPost, middlewares.MustProvideAccessToken)
	// Update the amount of some product in the cart
	e.PUT("/api/cart", controllers.HandleCartPut, middlewares.MustProvideAccessToken)

	//Delete cart
	e.DELETE("/api/cart", controllers.DeleteCartProduct, middlewares.MustProvideAccessToken)
}
