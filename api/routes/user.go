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
	// obtain favs
	e.GET("/api/user/favorites/list", controllers.FavoritesGET, middlewares.MustProvideAccessToken)
	// obtain detail favs
	e.GET("/api/user/favorites/detailed", controllers.HandleDetailedFavorites, middlewares.MustProvideAccessToken)
	// deleting favs
	e.DELETE("/api/user/favorites/:id", controllers.HandleDeletingFavorite, middlewares.MustProvideAccessToken)
	// obtain user orders
	e.GET("/api/user/orders", controllers.HandleGetOrders, middlewares.MustProvideAccessToken)
}
