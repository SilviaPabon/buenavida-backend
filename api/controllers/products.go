package controllers

import(
  "strconv"
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/SilviaPabon/buenavida-backend/interfaces"
  "github.com/SilviaPabon/buenavida-backend/models"
)

// HandleProductsPagination get products by given page
func HandleProductsPagination(c echo.Context) error {
  // Get page from params and convert to int
  param := c.Param("page")
  page, err := strconv.Atoi(param)

  if page <= 0 || err != nil{
    return c.JSON(http.StatusBadRequest, interfaces.GenericResponse{
      Error: true, 
      Message: "Page must be a possitive integer (Starting from zero)",
    })
  }

  // Get products by given page
  products, err := models.GetProductsByPage(page)

  if err != nil{
    return c.JSON(http.StatusInternalServerError, interfaces.GenericResponse{
      Error: true, 
      Message: "Unable to get products from database",
    })
  }

  if len(products) == 0{
    return c.JSON(http.StatusBadRequest, interfaces.GenericResponse{
      Error: true, 
      Message: "Found empty page", 
    })
  }

  return c.JSON(http.StatusOK, interfaces.ProductsPage{
    Error: false, 
    Message: "OK", 
    Products: products, 
  })

}
