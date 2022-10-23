package main

import(
  "fmt"
  "os"
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/SilviaPabon/buenavida-backend/configs"
)

func main(){
  // ### ### ###
  // Create database connection
  configs.ConnectToMongo()

  // ### ### ###
  // Echo setup
  e := echo.New()

  e.GET("/ping", func(c echo.Context) error {
    return c.String(http.StatusOK, "Pong!!")
  })

  // ### ### ###
  // Configure port

  port := os.Getenv("PORT")

  if port == "" {
    port = "3030"
  }

  e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
