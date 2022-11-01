package middlewares

import(
  "fmt"
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/SilviaPabon/buenavida-backend/interfaces"
)

// MustProvideAccessToken Verify access token is present as a cookie
func MustProvideAccessToken(next echo.HandlerFunc) echo.HandlerFunc {
  return func(c echo.Context) error {
    // Get token
    cookie, err := c.Cookie("access-token")

    if err != nil {
      return c.JSON(http.StatusUnauthorized, interfaces.GenericResponse{
	Error: true, 
	Message: "Access token wasn't provided",
      })
    }

    return next(c)
  }
} 
