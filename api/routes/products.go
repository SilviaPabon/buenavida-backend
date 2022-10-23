package routes

import(
  "github.com/labstack/echo/v4"
  "github.com/SilviaPabon/buenavida-backend/controllers"
)

// SetupProductsRoutes initialize user routes
func SetupProductsRoutes(e *echo.Echo){
  e.GET("/products/:page", controllers.HandleProductsPagination)
}
