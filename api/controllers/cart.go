package controllers

import(
  // "fmt"
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/golang-jwt/jwt/v4"
  "github.com/SilviaPabon/buenavida-backend/configs"
  "github.com/SilviaPabon/buenavida-backend/models"
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
  _, err := models.GetDetailsFromID(payload.Id.Hex())

  if err != nil {
    return c.JSON(http.StatusNotFound, interfaces.GenericResponse{
      Error: true, 
      Message: "Product was not found.",
    })
  }

  // Get user id from token
  cookie, _ := c.Cookie("access-token")
  token := cookie.Value
  claims := &interfaces.JWTCustomClaims{}
  jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error){
    return configs.GetJWTSecret(), nil
  })

  // Save on database
  err = models.AddProductToCart(claims.ID, payload.Id.Hex())

  if err != nil {
    return c.JSON(http.StatusInternalServerError, interfaces.GenericResponse{
      Error: true, 
      Message: "Unable to add the product to the user cart. Try again.",
    })
  }

  return c.JSON(http.StatusOK, interfaces.GenericResponse{
    Error: false, 
    Message: "Product added to the cart successfully",
  })
}

// HandleCartPut Update the amount of some product in the cart
func HandleCartPut(c echo.Context) error {
  return c.JSON(http.StatusOK, interfaces.GenericResponse{
    Error: false, 
    Message: "Received",
  })
}
