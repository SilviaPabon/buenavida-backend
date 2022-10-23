package main

import(
  "github.com/labstack/echo/v4"
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

  // Start routes
  routes.SetupProductsRoutes(e)

  e.Logger.Fatal(e.Start(":3030"))
}
