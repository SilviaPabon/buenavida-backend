package controllers

import(
  "fmt"
  "testing"
  "net/http"
  "encoding/json"
  "net/http/httptest"
  "github.com/labstack/echo/v4"
  "github.com/stretchr/testify/require"
  "github.com/SilviaPabon/buenavida-backend/interfaces"
)

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
  products := len(reply.Products)
  c.GreaterOrEqualf(products, 25, fmt.Sprintf("Got: %d videos --> At least 25 videos are required", products))
}
