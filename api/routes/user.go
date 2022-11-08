package routes

import (
	"github.com/SilviaPabon/buenavida-backend/controllers"
	"github.com/SilviaPabon/buenavida-backend/middlewares"
	"github.com/labstack/echo/v4"
)

// SetupuserRoutes create and manage user routes
func SetupUserRoutes(e *echo.Echo) {
	e.POST("/api/user", controllers.HandleUserPost)
	// update favorites from user
	e.POST("/api/user/favorites", controllers.FavoritesPost, middlewares.MustProvideAccessToken)
	e.GET("/api/user/favorites/list", controllers.FavoritesGET, middlewares.MustProvideAccessToken)
	e.GET("/api/user/favorites/detailed", controllers.HandleDetailedFavorites, middlewares.MustProvideAccessToken)
}
