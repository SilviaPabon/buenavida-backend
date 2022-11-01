package routes

import (
	"github.com/SilviaPabon/buenavida-backend/controllers"
	"github.com/labstack/echo/v4"
)

// SetupuserRoutes create and manage user routes
func SetupUserRoutes(e *echo.Echo) {

	/* e.GET("/user/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Pong!!")
	}) */

	//e.GET("/api/user/:id", controllers.HandleUserGet)
	e.POST("/api/user", controllers.HandleUserPost)
}
