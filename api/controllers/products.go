package controllers

import(
  "strconv"
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/SilviaPabon/buenavida-backend/interfaces"
  //"github.com/SilviaPabon/buenavida-backend/models"
)

// HandleProductsPagination get products by given page
func HandleProductsPagination(c echo.Context) error {
  // Get page from params
  param := c.Param("page")
  _, err := strconv.Atoi(param)

  if err != nil {
    return c.JSON(http.StatusBadRequest, interfaces.GenericResponse{
      Error: true, 
      Message: "Page must be a numeric value",
    })
  }

  // Get products by given page
  // models.GetProducsByPage(id)

  return c.JSON(http.StatusOK, map[string]string{
    "Message": "Products route",
  })
}
