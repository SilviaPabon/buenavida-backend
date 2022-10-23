package controllers

import(
  "net/http"
  "github.com/labstack/echo/v4"
)

// HandleProductsPagination get products by given page
func HandleProductsPagination(c echo.Context) error {
  return c.JSON(http.StatusOK, map[string]string{
    "Message": "Products route",
  })
}
