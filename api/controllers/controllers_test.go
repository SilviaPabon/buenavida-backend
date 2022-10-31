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

// Helper function to create the context and recorder
func setup(method, path string) (echo.Context, *httptest.ResponseRecorder, *http.Request) {
  e := echo.New()
  r := httptest.NewRequest(method, path, nil)
  w := httptest.NewRecorder()
  context := e.NewContext(r, w)

  return context, w, r
}

// Test /api/producst success
func TestGetProductsSuccess(t *testing.T){
  c := require.New(t)

  // Create request
  context, w, _ := setup(http.MethodGet, "/api/products")

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

  c.GreaterOrEqualf(len(reply.Products), 25, fmt.Sprintf("Got: %d products --> At least 25 products are required", len(reply.Products)))
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
  context, w, _ := setup(http.MethodGet, "/api/products")

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

// Test /api/products/1 success
func TestProductsPaginationSuccess(t *testing.T){
  c := require.New(t)

  context, w, _ := setup(http.MethodGet, "/api/products")
  context.SetParamNames("page")
  context.SetParamValues("1")

  err := HandleProductsPagination(context)
  c.NoError(err)

  // Validate custom request responses
  c.Equal(http.StatusOK, w.Code)
  
  var reply interfaces.GenericProductsArrayResponse
  err =  json.Unmarshal(w.Body.Bytes(), &reply)
  c.NoError(err)

  c.Equal(reply.Error, false)
  c.Equalf(len(reply.Products), 12, fmt.Sprintf("Got: %d products --> Expected pagination was 12 items", len(reply.Products)))
}
