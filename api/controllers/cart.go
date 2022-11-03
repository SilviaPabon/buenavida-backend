package controllers

import(
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/SilviaPabon/buenavida-backend/interfaces"
)

//  HandleCartPost add a new product to the cart
func HandleCartPost(c echo.Context) error {
  return c.JSON(http.StatusOK, interfaces.GenericResponse{
    Error: false, 
    Message: "OK",
  })
}
