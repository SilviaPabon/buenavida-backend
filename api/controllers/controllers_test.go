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

// #### #### #### #### ####
// #### #### Products #### ####
// #### #### #### #### ####

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
  c.Equalf(http.StatusOK, w.Code, fmt.Sprintf("Expected %d status code but got %d", http.StatusOK, w.Code))

  // Convert response to struct
  var reply interfaces.GenericProductsArrayResponse
  err = json.Unmarshal(w.Body.Bytes(), &reply)
  c.NoError(err)

  // Validate response
  c.Equalf(false, reply.Error, fmt.Sprintf("Expected custom error on JSON to be false but got %t", reply.Error))

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
  c.Equalf(http.StatusInternalServerError, w.Code, fmt.Sprintf("Expected %d status code but got %d", http.StatusInternalServerError, w.Code))
 
  // Validate body (Just the error boolean because there are not products)
  var reply interfaces.GenericProductsArrayResponse
  err = json.Unmarshal(w.Body.Bytes(), &reply)
  c.NoError(err)

  c.Equalf(true, reply.Error, fmt.Sprintf("Expected custom error on JSON to be true but got %t", reply.Error))
}

// Test /api/products/:id success
func TestProductsPaginationSuccess(t *testing.T){
  c := require.New(t)

  context, w, _ := setup(http.MethodGet, "/api/products")
  context.SetParamNames("page")
  context.SetParamValues("1")

  err := HandleProductsPagination(context)
  c.NoError(err)

  // Validate custom request responses
  c.Equalf(http.StatusOK, w.Code, fmt.Sprintf("Expected %d status code but got %d", http.StatusOK, w.Code))
  
  var reply interfaces.GenericProductsArrayResponse
  err = json.Unmarshal(w.Body.Bytes(), &reply)
  c.NoError(err)

  c.Equal(false, reply.Error)
  c.Equalf(12, len(reply.Products), fmt.Sprintf("Got: %d products --> Expected pagination was 12 items", len(reply.Products)))
}

// Test /api/products/:id Not found
func TestProductsPaginationNotFound(t *testing.T){
  c := require.New(t)

  context, w, _ := setup(http.MethodGet, "/api/products")
  context.SetParamNames("page")
  context.SetParamValues("939")

  err := HandleProductsPagination(context)
  c.NoError(err)

  // Validate custom request responses
  c.Equalf(http.StatusNotFound, w.Code, fmt.Sprintf("Expected %d status code but got %d", http.StatusNotFound, w.Code))

  var reply interfaces.GenericResponse
  err = json.Unmarshal(w.Body.Bytes(), &reply)
  c.NoError(err)

  c.Equalf(true, reply.Error, fmt.Sprintf("Expected custom error on JSON to be true but got %t", reply.Error))
}

// Test /api/products/:id Bad Request
func TestProductsPaginationBadRequest(t *testing.T){
  c := require.New(t)

  context, w, _ := setup(http.MethodGet, "/api/products")

  // Test with not valid param
  context.SetParamNames("page")
  context.SetParamValues("-1")

  err := HandleProductsPagination(context)
  c.NoError(err)

  // Validate custom request responses
  c.Equalf(http.StatusBadRequest, w.Code, fmt.Sprintf("Expected %d status code but got %d", http.StatusBadRequest, w.Code))

  // Test with not given param
  context, w, _ = setup(http.MethodGet, "/api/products")
  err = HandleProductsPagination(context)
  c.NoError(err)

  c.Equalf(http.StatusBadRequest, w.Code, fmt.Sprintf("Expected %d status code but got %d", http.StatusBadRequest, w.Code))
 
  // Validate custom error field
  var reply interfaces.GenericResponse
  err = json.Unmarshal(w.Body.Bytes(), &reply)
  c.NoError(err)

  c.Equalf(true, reply.Error, fmt.Sprintf("Expected custom error on JSON to be true but got %t", reply.Error))
}

// Test /api/products/:id Internal server error
func TestProductsPaginationInternalServerError(t *testing.T){
  c := require.New(t)
  context, w, _ := setup(http.MethodGet, "/api/products")
  context.SetParamNames("page")
  context.SetParamValues("1")

  // Mock models method
  originalGetProductsByPage := modelsGetProductsByPage
  defer func(){ modelsGetProductsByPage = originalGetProductsByPage }()

  modelsGetProductsByPage = func(page int) ([]interfaces.Article, error) {
    // Return intentional error
    return []interfaces.Article{}, errors.New("Oops...")
  }

  // Send request
  err := HandleProductsPagination(context)
  c.NoError(err)

  c.Equalf(http.StatusInternalServerError, w.Code, fmt.Sprintf("Extected %d status code but got %d", http.StatusInternalServerError, w.Code))

  var reply interfaces.GenericResponse
  err = json.Unmarshal(w.Body.Bytes(), &reply)
  c.NoError(err)

  c.Equalf(true, reply.Error, fmt.Sprintf("Expected custom error on JSON to be true but got %t", reply.Error))
}

// #### #### #### #### ####
// #### #### User #### ####
// #### #### #### #### ####

// To do
