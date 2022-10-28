package routes

import(
  "github.com/labstack/echo/v4"
  "github.com/SilviaPabon/buenavida-backend/controllers"
)

// SetupProductsRoutes initialize user routes
func SetupProductsRoutes(e *echo.Echo){
  // Get all products
  e.GET("/api/products", controllers.HandleProductsGet)
  // Get products by page
  e.GET("/api/products/:page", controllers.HandleProductsPagination)
  // Search product fron text
  e.POST("/api/products/search", controllers.HandleProductsSearch)
  // Get product image
  e.GET("/api/products/image/:serial", controllers.HandleProductImageRequest)
}
