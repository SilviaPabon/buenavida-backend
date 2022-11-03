package controllers

import(
  //"fmt"
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/SilviaPabon/buenavida-backend/interfaces"
)

//  HandleCartPost add a new product to the cart
func HandleCartPost(c echo.Context) error {
  // Get json payload
  payload := new(interfaces.AddToCartPayload)

  if err := c.Bind(payload); err != nil {
    return c.JSON(http.StatusBadRequest, interfaces.GenericResponse{
      Error: true, 
      Message: "Unable to process request.",
    })
  }

  if payload.Id.IsZero() {
    return c.JSON(http.StatusBadRequest, interfaces.GenericResponse{
      Error: true, 
      Message: "Provided object id is emtpy or not valid",
    })
  }

  // Validate the product exists on mongo

  return c.JSON(http.StatusOK, interfaces.GenericResponse{
    Error: false, 
    Message: "OK",
  })
}
