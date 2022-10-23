package main

import(
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "github.com/SilviaPabon/buenavida-backend/configs"
  "github.com/SilviaPabon/buenavida-backend/routes"
)

func main(){
  // ### ### ###
  // Create database connection
  configs.ConnectToMongo()

  // ### ### ###
  // Echo setup
  e := echo.New()

  e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
      AllowOrigins: []string{"*"},
      AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
  }))

  // Start routes
  routes.SetupProductsRoutes(e)

  e.Logger.Fatal(e.Start(":3030"))
}
