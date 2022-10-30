package controllers

import(
  "fmt"
  "strconv"
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/SilviaPabon/buenavida-backend/interfaces"
  "go.mongodb.org/mongo-driver/bson/primitive"
  "github.com/SilviaPabon/buenavida-backend/models"
)

// HandleProductsGet
func HandleProductsGet(c echo.Context) error {
  // Get all products
  products, err := models.GetAllProducts()

  if err != nil {
    return c.JSON(http.StatusInternalServerError, interfaces.GenericResponse{
      Error:   true,
      Message: "Unable to get products from database. Try again.",
    })
  }

  return c.JSON(http.StatusOK, interfaces.GenericProductsArrayResponse{
    Error:    false,
    Message:  "OK",
    Products: products,
  })

}

// HandleProductsPagination get products by given page
func HandleProductsPagination(c echo.Context) error {
  // Get page from params and convert to int
  param := c.Param("page")
  page, err := strconv.Atoi(param)
  println(param)

  if page <= 0 || err != nil {
    return c.JSON(http.StatusBadRequest, interfaces.GenericResponse{
      Error:   true,
      Message: "Page must be a possitive integer (Starting from zero)",
    })
  }

  // Get products by given page
  products, err := models.GetProductsByPage(page)

  if err != nil {
    return c.JSON(http.StatusInternalServerError, interfaces.GenericResponse{
      Error:   true,
      Message: "Unable to get products from database. Try again.",
    })
  }

  if len(products) == 0 {
    return c.JSON(http.StatusNotFound, interfaces.GenericResponse{
      Error:   true,
      Message: "Page wasn't found",
    })
  }

  return c.JSON(http.StatusOK, interfaces.GenericProductsArrayResponse{
    Error:    false,
    Message:  "OK",
    Products: products,
  })

}

// HandleProductsSearch
func HandleProductsSearch(c echo.Context) error {
  // Get json payload
  payload := new(interfaces.FilterProductsByText)

  if err := c.Bind(payload); err != nil {
    return c.JSON(http.StatusBadRequest, interfaces.GenericResponse{
      Error:   true,
      Message: "Unable to process query. Try again and make sure search-criteria field is provided",
    })
  }

  // Validate field is not empty
  if payload.Criteria == "" {
    return c.JSON(http.StatusBadRequest, interfaces.GenericResponse{
      Error:   true,
      Message: "search-criteria must be provided and can't be empty",
    })
  }

  // Search in database
  products, dberr := models.SearchByText(payload.Criteria)

  if dberr != nil {
    return c.JSON(http.StatusInternalServerError, interfaces.GenericResponse{
      Error:   true,
      Message: "Unable to get response fron database",
    })
  }

  return c.JSON(http.StatusOK, interfaces.GenericProductsArrayResponse{
    Error:    false,
    Message:  "OK",
    Products: products,
  })
}

func GetFromID(c echo.Context) error {
  //param id
  id := c.Param("id")

  product, err := models.GetDetailsFromID(id)
  fmt.Println(err)

  if err != nil {
    switch err {
    case primitive.ErrInvalidHex:
      return c.JSON(http.StatusBadRequest, interfaces.GenericResponse{
	Error:   true,
	Message: "Provided object id is not valid",
      })
    default:
      return c.JSON(http.StatusNotFound, interfaces.GenericResponse{
	Error:   true,
	Message: "Unable to get products from database. Try again.",
      })
    }
  }

  return c.JSON(http.StatusOK, interfaces.GenericProductResponse{
    Error:   false,
    Message: "OK",
    Product: product,
  })
}

// HandleProductImageRequest get product image fron given serial
func HandleProductImageRequest(c echo.Context) error {
  // Validate serial is a number
  serialString := c.Param("serial")
  serialNumber, err := strconv.Atoi(serialString)

  if err != nil {
    return c.JSON(http.StatusBadRequest, interfaces.GenericResponse{
      Error: true, 
      Message: "Provided serial is not valid. You must provide an positive number",
    })
  }

  image, err := models.GetProductImageFromSerial(serialNumber)

  if err != nil {
    return c.JSON(http.StatusNotFound, interfaces.GenericResponse{
      Error: true, 
      Message: fmt.Sprintf("No image was found for product with id %d", serialNumber),
    })
  }

  return c.JSON(http.StatusOK, interfaces.ProductImageResponse{
    Error: false, 
    Message: "OK", 
    Image: image.Image,
  })
}
