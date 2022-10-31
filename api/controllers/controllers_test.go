package controllers

import(
  "fmt"
  "errors"
  "testing"
  "net/http"
  "encoding/json"
  "net/http/httptest"
  "github.com/labstack/echo/v4"
  "github.com/stretchr/testify/require"
  "github.com/SilviaPabon/buenavida-backend/interfaces"
)

// Test /api/producst success
func TestGetProductsSuccess(t *testing.T){
  c := require.New(t)

  // Create request
  e := echo.New()
  r := httptest.NewRequest(http.MethodGet, "/api/products", nil)
  w := httptest.NewRecorder()
  context := e.NewContext(r, w)

  // Make request
  err := HandleProductsGet(context)
  c.NoError(err)

  // Validate status code
  c.Equal(w.Code, http.StatusOK)

  // Convert response to struct
  var reply interfaces.GenericProductsArrayResponse
  err = json.Unmarshal(w.Body.Bytes(), &reply)
  c.NoError(err)

  // Validate response
  c.Equal(reply.Error, false)

  c.GreaterOrEqualf(len(reply.Products), 25, fmt.Sprintf("Got: %d videos --> At least 25 videos are required", len(reply.Products)))
}

// Test /api/products failed
func TestGetProductsInternalServerError(t *testing.T) {
  c := require.New(t)

  // Save original functions references
  originalGetAllProducts := modelsGetAllProducts
  defer func(){ modelsGetAllProducts = originalGetAllProducts }()

  // Mock function
  modelsGetAllProducts = func() ([]interfaces.Article, error) {
    // Return intentional error
    return []interfaces.Article{}, errors.New("Oops...")
  }

  // Create request
  e := echo.New()
  r := httptest.NewRequest(http.MethodGet, "/api/products", nil)
  w := httptest.NewRecorder()
  context := e.NewContext(r, w)

  // Make request
  err := HandleProductsGet(context)
  c.NoError(err)

  // Validate status code
  c.Equal(w.Code, http.StatusInternalServerError)
 
  // Validate body (Just the error boolean because there are not products)
  var reply interfaces.GenericProductsArrayResponse
  err = json.Unmarshal(w.Body.Bytes(), &reply)
  c.NoError(err)

  c.Equal(reply.Error, true)
}
