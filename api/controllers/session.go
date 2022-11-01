package controllers

import(
  "fmt"
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/SilviaPabon/buenavida-backend/interfaces"
  "github.com/SilviaPabon/buenavida-backend/models"
  "github.com/SilviaPabon/buenavida-backend/utils"
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

  // Get user from database
  user, err := models.GetUserFromMail(payload.Mail)

  if err != nil {
    return c.JSON(http.StatusNotFound, interfaces.GenericResponse{
      Error: true, 
      Message: "User wasn't found",
    })
  }

  // fmt.Printf("%+v\n", user)
  accessToken, ATerr := utils.CreateJWTAccessToken(&user)
  refreshToken, RTerr := utils.CreateJWTRefreshToken(&user)

  if ATerr != nil || RTerr != nil {
    return c.JSON(http.StatusInternalServerError, interfaces.GenericResponse{
      Error: true, 
      Message: "Unable to initialize authentication",
    })
  }

  fmt.Println(accessToken)
  fmt.Println(refreshToken)

  return c.JSON(http.StatusOK, interfaces.GenericResponse{
    Error: false, 
    Message: "Hello world",
  })
}
