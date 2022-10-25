package main

import(
  "fmt"
  "os"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "github.com/SilviaPabon/buenavida-backend/configs"
  "github.com/SilviaPabon/buenavida-backend/routes"
)

func main(){
  // ### ### ###
  // Create mongo database connection
  configs.ConnectToMongo()
 
  // Testing postgres connection (THIS SHOULD BE DELETED IN FUTURE)
  db := configs.ConnectToPostgres()
  pgPingErr := db.Ping()
  
  if pgPingErr != nil {
    panic("🟥 Unable to ping postgres database 🟥")
  }else{
    fmt.Println("🐘 Connected to postgresSQL")
  }

  defer db.Close()

  // ### ### ###
  // Echo setup
  e := echo.New()

  e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
      AllowOrigins: []string{"*"},
      AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
  }))

  // Start routes
  routes.SetupProductsRoutes(e)

  // ### ### ###
  // Configure port

  port := os.Getenv("PORT")

  if port == "" {
    port = "3030"
  }

  e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
