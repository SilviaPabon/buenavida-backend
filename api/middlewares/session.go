package middlewares

import(
  // "fmt"
  // "time"
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/golang-jwt/jwt/v4"
  "github.com/SilviaPabon/buenavida-backend/configs"
  "github.com/SilviaPabon/buenavida-backend/interfaces"
  "github.com/SilviaPabon/buenavida-backend/models"
)

// MustProvideAccessToken Verify access token is present as a cookie
func MustProvideAccessToken(next echo.HandlerFunc) echo.HandlerFunc {
  return func(c echo.Context) error {
    // Get token
    cookie, err := c.Cookie("access-token")

    if err != nil {
      return c.JSON(http.StatusForbidden, interfaces.GenericResponse{
	Error: true, 
	Message: "Access token wasn't provided",
      })
    }

    claims := &interfaces.JWTCustomClaims{}
    signedString := cookie.Value

    // This also validates the token expiration
    _, err = jwt.ParseWithClaims(signedString, claims, func(t *jwt.Token) (interface {}, error){
      return configs.GetJWTSecret(), nil
    })
    
    if err != nil {
      return c.JSON(http.StatusForbidden, interfaces.GenericResponse{
	Error: true, 
	Message: "Access token is not valid",
      })
    }

    // Go to the handler if all is valid
    return next(c)
  }
}

// MustProvideRefreshToken Verify refresh token is present as a cookie
func MustProvideRefreshToken(next echo.HandlerFunc) echo.HandlerFunc {
  return func(c echo.Context) error {
    // Get token
    cookie, err := c.Cookie("refresh-token")

    if err != nil {
      return c.JSON(http.StatusForbidden, interfaces.GenericResponse{
	Error: true, 
	Message: "Refresh token was not provided",
      })
    }

    // Validate claims
    claims := &interfaces.JWTCustomClaims{}
    signedString := cookie.Value

    _, err = jwt.ParseWithClaims(signedString, claims, func(t *jwt.Token) (interface {}, error){
      return configs.GetJWTSecret(), nil
    })
    
    if err != nil {
      return c.JSON(http.StatusForbidden, interfaces.GenericResponse{
	Error: true, 
	Message: "Refresh token is not valid",
      })
    }
    
    // Verify token exists on redis white-list
    storedToken, err := models.GetRefreshTokenFromRedis(claims.Email)

    if err != nil || storedToken != claims.UUID.String() {
      return c.JSON(http.StatusForbidden, interfaces.GenericResponse{
	Error: true, 
	Message: "Refresh token is not authentic",
      })
    }

    return next(c)
    
  }
}
