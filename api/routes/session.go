package routes

import(
  "github.com/labstack/echo/v4"
  "github.com/SilviaPabon/buenavida-backend/controllers"
  // "github.com/SilviaPabon/buenavida-backend/middlewares"
)

// SetupSessionRoutes create and manage session routes
func SetupSessionRoutes(e *echo.Echo){
  e.POST("/api/session/login", controllers.HandleLogin)
}
