package controllers

import(
  "fmt"
  "time"
  "bytes"
  "context"
  "errors"
  "testing"
  "net/http"
  "encoding/json"
  "net/http/httptest"
  "github.com/labstack/echo/v4"
  "github.com/go-faker/faker/v4"
  "github.com/stretchr/testify/require"
  "github.com/SilviaPabon/buenavida-backend/configs" // db connection
  "github.com/SilviaPabon/buenavida-backend/interfaces"
)

var pg = configs.ConnectToPostgres()

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

// Helper function to create a post request
func setupPost(method, path string, payload interface{}) (echo.Context, *httptest.ResponseRecorder, *http.Request){
  payloadBytes, _ := json.Marshal(payload)

  e := echo.New()
  r := httptest.NewRequest(method, path, bytes.NewBuffer(payloadBytes))
  r.Header.Set("Content-Type", "application/json")
  w := httptest.NewRecorder()
  context := e.NewContext(r,w)

  return context, w, r
}

// TestGetProductsSuccess /api/producst success
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

// TestGetProductsInternalServerError /api/products failed
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

// TestProductsPaginationSuccess /api/products/:id success
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

// TestProductsPaginationNotFound /api/products/:id Not found
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

// TestProductsPaginationBadRequest /api/products/:id Bad Request
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

// TestProductsPaginationInternalServerError /api/products/:id Internal server error
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

// TestProductsSearchSuccess /api/products/search success
func TestProductsSearchSuccess(t *testing.T) {
  c := require.New(t)

  // Payload
  payload1 := interfaces.FilterProductsByText{
    Criteria: "Aceite",
  }

  // Make request 1
  context, w, _ := setupPost(http.MethodPost, "/api/products/search", payload1)
  err := HandleProductsSearch(context)
  c.NoError(err)

  var reply interfaces.GenericProductsArrayResponse
  err = json.Unmarshal(w.Body.Bytes(), &reply)
  c.NoError(err)

  c.GreaterOrEqual(len(reply.Products), 21, fmt.Sprintf("Exptected at least %d products matching the search criteria but gott %d", 21, len(reply.Products)))
  c.Equalf(false, reply.Error, fmt.Sprintf("Exptected cuscom error to be false but found: %t", reply.Error))

  // Make request 2
  payload2 := interfaces.FilterProductsByText{
    Criteria: "aCeIte",
  }

  // Make request 1
  context, w, _ = setupPost(http.MethodPost, "/api/products/search", payload2)
  err = HandleProductsSearch(context)
  c.NoError(err)

  var reply2 interfaces.GenericProductsArrayResponse
  err = json.Unmarshal(w.Body.Bytes(), &reply2)
  c.NoError(err)

  c.GreaterOrEqual(len(reply2.Products), 21, fmt.Sprintf("Exptected at least %d products matching the search criteria but gott %d", 21, len(reply2.Products)))
  c.Equalf(false, reply2.Error, fmt.Sprintf("Exptected custom error to be false but found: %t", reply2.Error))
}

// TestProductDetailsSuccess /api/product/:id 
func TestProductDetailsSucess(t *testing.T){
  c := require.New(t)

  context, w, _ := setup(http.MethodGet, "/api/product")
  context.SetParamNames("id")

  // IMPORTANT: This id can change if you run the bulkdata command again
  // So, replace it with a valid MongoID in your case
  targetId := "635f406d344c343aabfee5f1"
  context.SetParamValues(targetId)

  // Make request
  var reply interfaces.GenericProductResponse

  err := GetFromID(context)
  c.NoError(err)

  err = json.Unmarshal(w.Body.Bytes(), &reply)
  c.NoError(err)

  c.Equalf(http.StatusOK, w.Code, fmt.Sprintf("Exptected status code to be: %d but got: %d", http.StatusOK, w.Code))
  c.Equalf(false, reply.Error, fmt.Sprintf("Exptected custom error to be false but got: %t", reply.Error))
  c.Equalf(targetId, reply.Product.ID.Hex(), fmt.Sprintf("Expected result product id: %s to be equal than target id: %s", reply.Product.ID.Hex(), targetId))

  // Vefiry non-zero values
  c.NotEqual("", reply.Product.Name) 
  c.NotEqual("", reply.Product.Image) 
  c.NotEqual(0.0, reply.Product.Price) 
}

// #### #### #### #### ####
// #### #### User #### ####
// #### #### #### #### ####

// TestSignupSuccess /api/user POST
func TestSignupSuccess(t *testing.T){
  c := require.New(t)

  ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
  defer cancel()

  // Payload
  randEmail := faker.Email() // Save email to delete the user
  payload := interfaces.User{
    Firstname: faker.FirstName(),
    Lastname: faker.LastName(), 
    Email: randEmail, 
    Password: faker.Password() + "#1",  
  }

  // Make request
  context, w, _ := setupPost(http.MethodPost, "/api/user", payload)
  err := HandleUserPost(context)
  c.NoError(err)

  // Validate response
  var reply interfaces.GenericResponse
  err = json.Unmarshal(w.Body.Bytes(), &reply)
  c.NoError(err)

  c.Equalf(http.StatusOK, w.Code, fmt.Sprintf("Expected status code to be: %d but got: %d", http.StatusOK, w.Code))
  c.Equalf(false, reply.Error, fmt.Sprintf("Expected custom error to be false but got: %t", reply.Error))

  // Remove user from database
  query := `DELETE FROM users WHERE "mail" = $1`
  pg.QueryRowContext(ctx, query, randEmail)
}

// TestSignupDuplicatedMail /api/user POST
func TestSignupDuplicatedMail(t *testing.T){
  c := require.New(t)

  ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
  defer cancel()
  
  randEmail := faker.Email()
  payload := interfaces.User{
    Firstname: faker.FirstName(),
    Lastname: faker.LastName(), 
    Email: randEmail, 
    Password: faker.Password() + "#1", 
  }

  // Make request (Save user first time)
  context, w, _ := setupPost(http.MethodPost, "/api/user", payload)
  err := HandleUserPost(context)
  c.NoError(err)
  c.Equal(http.StatusOK, w.Code)

  // Try to save the user again (Tey to save a duplicated user)
  context, w, _ = setupPost(http.MethodPost, "/api/user", payload)
  err = HandleUserPost(context)
  c.NoError(err)
  c.Equalf(http.StatusConflict, w.Code, fmt.Sprintf("Expected status code to be: %d due to duplicated email but found: %d", http.StatusConflict, w.Code))
  
  // Remove user from database
  query := `DELETE FROM users WHERE "mail" = $1`
  pg.QueryRowContext(ctx, query, randEmail)
}
