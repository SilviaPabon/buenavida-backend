package routes

import (
	"github.com/SilviaPabon/buenavida-backend/controllers"
	"github.com/SilviaPabon/buenavida-backend/middlewares"
	"github.com/labstack/echo/v4"
)

// SetupSessionRoutes create and manage session routes
func SetupSessionRoutes(e *echo.Echo) {
	e.POST("/api/session/login", controllers.HandleLogin)
	e.GET("/api/session/whoami", controllers.HandleWhoAmI, middlewares.MustProvideAccessToken)
	e.GET("/api/session/refresh", controllers.HandleRefresh, middlewares.MustProvideRefreshToken)
	e.DELETE("/api/session/logout", controllers.HandleLogout, middlewares.MustProvideAccessToken)
}
