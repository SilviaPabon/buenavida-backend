package routes

import (
	"github.com/SilviaPabon/buenavida-backend/controllers"
	"github.com/SilviaPabon/buenavida-backend/middlewares"
	"github.com/labstack/echo/v4"
)

// SetupuserRoutes create and manage user routes
func SetupUserRoutes(e *echo.Echo) {

	/* e.GET("/user/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Pong!!")
	}) */

	//e.GET("/api/user/:id", controllers.HandleUserGet)
	e.POST("/api/user", controllers.HandleUserPost)
	// update favorites from user
	e.POST("/api/user/favorites", controllers.FavoritesPost, middlewares.MustProvideAccessToken)
	// Obtain favorites from user
	e.GET("/api/user/favorites", controllers.FavoritesGET, middlewares.MustProvideAccessToken)
	// Obtain cart from user
	e.GET("/api/user/cart", controllers.HandleCartGet, middlewares.MustProvideAccessToken)
}
