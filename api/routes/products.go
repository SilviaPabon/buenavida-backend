package routes

import(
  "github.com/labstack/echo/v4"
  "github.com/SilviaPabon/buenavida-backend/controllers"
)

// SetupProductsRoutes initialize user routes
func SetupProductsRoutes(e *echo.Echo){
  // Get all products
  e.GET("/products", controllers.HandleProductsGet)
  // Get products by page
  e.GET("/products/:page", controllers.HandleProductsPagination)
  // Search product fron text
  e.POST("/products/search", controllers.HandleProductsSearch)
}
