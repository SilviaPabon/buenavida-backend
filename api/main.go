package main

import(
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

  e.Logger.Fatal(e.Start(":3030"))
}
