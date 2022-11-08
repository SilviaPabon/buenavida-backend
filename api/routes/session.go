package routes

import (
	"github.com/SilviaPabon/buenavida-backend/controllers"
	"github.com/SilviaPabon/buenavida-backend/middlewares"
	"github.com/labstack/echo/v4"
)

// SetupSessionRoutes create and manage session routes
func SetupSessionRoutes(e *echo.Echo) {
	e.POST("/api/session/login", controllers.HandleLogin)

	// This route should be deleted, is just for testing
	e.GET("/api/session/ping", controllers.HandlePing, middlewares.MustProvideAccessToken)

	// Uncoment this to test the refresh token
	// e.GET("/api/session/refresh", controllers.HandlePing, middlewares.MustProvideAccessToken, middlewares.MustProvideRefreshToken)

	//Logout
	e.POST("/api/session/logout", controllers.HandleLogout, middlewares.MustProvideAccessToken)
}
