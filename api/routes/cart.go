package routes

import(
  "github.com/labstack/echo/v4"
  "github.com/SilviaPabon/buenavida-backend/controllers"
  "github.com/SilviaPabon/buenavida-backend/middlewares"
)

// SetupCartRoutes create an manage cart routes
func SetupCartRoutes(e *echo.Echo){
  // Add a new product to the cart
  e.POST("/api/cart", controllers.HandleCartPost, middlewares.MustProvideAccessToken)
}
