package controllers

import(
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/SilviaPabon/buenavida-backend/interfaces"
)

func HandleLogin(c echo.Context) error {
  // Get json payload
  var payload = new(interfaces.LoginPayload)
  err := c.Bind(payload)

  if err != nil {
    return c.JSON(http.StatusBadRequest, interfaces.GenericResponse{
      Error: true, 
      Message: "Unable to process query. Try again and make sure mail and password fields were provided",
    })
  }

  // Validate json payload
  if payload.Mail == "" || payload.Password == "" {
    return c.JSON(http.StatusBadRequest, interfaces.GenericResponse{
      Error: true, 
      Message: "Mail and password files are required / can't be empty",
    })
  }

  return c.JSON(http.StatusOK, interfaces.GenericResponse{
    Error: false, 
    Message: "Hello world",
  })
}
