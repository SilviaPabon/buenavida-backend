package controllers

import(
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/SilviaPabon/buenavida-backend/interfaces"
)

func HandleLogin(c echo.Context) error {
  return c.JSON(http.StatusOK, interfaces.GenericResponse{
    Error: false, 
    Message: "Hello world",
  })
}
