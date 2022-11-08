package controllers

import (
	// "fmt"

	"net/http"

	"github.com/SilviaPabon/buenavida-backend/configs"
	"github.com/SilviaPabon/buenavida-backend/interfaces"
	"github.com/SilviaPabon/buenavida-backend/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// HandleCartPost add a new product to the cart
func HandleCartPost(c echo.Context) error {
	// Get json payload
	payload := new(interfaces.ProductIdPayload)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, interfaces.GenericResponse{
			Error:   true,
			Message: "Unable to process request.",
		})
	}

	if payload.Id.IsZero() {
		return c.JSON(http.StatusBadRequest, interfaces.GenericResponse{
			Error:   true,
			Message: "Provided object id is emtpy or not valid",
		})
	}

	// Validate the product exists on mongo
	_, err := models.GetDetailsFromID(payload.Id.Hex())

	if err != nil {
		return c.JSON(http.StatusNotFound, interfaces.GenericResponse{
			Error:   true,
			Message: "Product was not found.",
		})
	}

	// Get user id from token
	cookie, _ := c.Cookie("access-token")
	token := cookie.Value
	claims := &interfaces.JWTCustomClaims{}
	jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return configs.GetJWTSecret(), nil
	})

	// Save on database
	err = models.AddProductToCart(claims.ID, payload.Id.Hex())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, interfaces.GenericResponse{
			Error:   true,
			Message: "Unable to add the product to the user cart. Try again.",
		})
	}

	return c.JSON(http.StatusOK, interfaces.GenericResponse{
		Error:   false,
		Message: "Product added to the cart successfully",
	})
}

// HandleCartPut Update the amount of some product in the cart
func HandleCartPut(c echo.Context) error {
	// Get json payload
	payload := new(interfaces.UpdateCartPayload)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, interfaces.GenericResponse{
			Error:   true,
			Message: "Unable to process request.",
		})
	}

	if payload.Id.IsZero() { // Validate provided object id
		return c.JSON(http.StatusBadRequest, interfaces.GenericResponse{
			Error:   true,
			Message: "Provided object id is emtpy or not valid",
		})
	}

	if payload.Amount <= 0 || payload.Amount >= 128 { // Validate amount
		return c.JSON(http.StatusBadRequest, interfaces.GenericResponse{
			Error:   true,
			Message: "Product amount must be greater than zero and lower than 128.",
		})
	}

	// Validate the product exists on mongo
	_, err := models.GetDetailsFromID(payload.Id.Hex())

	if err != nil {
		return c.JSON(http.StatusNotFound, interfaces.GenericResponse{
			Error:   true,
			Message: "Product was not found.",
		})
	}

	// Get user id from token
	cookie, _ := c.Cookie("access-token")
	token := cookie.Value
	claims := &interfaces.JWTCustomClaims{}
	jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return configs.GetJWTSecret(), nil
	})

	// Make database query
	err = models.UpdateProductInCart(claims.ID, payload.Amount, payload.Id.Hex())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, interfaces.GenericResponse{
			Error:   true,
			Message: "Unable to update product in cart. Try again.",
		})
	}

	return c.JSON(http.StatusOK, interfaces.GenericResponse{
		Error:   false,
		Message: "Cart updated successfully",
	})
}

func DeleteCartProduct(c echo.Context) error {

	payload := new(interfaces.ProductIdPayload)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, interfaces.GenericResponse{
			Error:   true,
			Message: "Unable to process request.",
		})
	}

	if payload.Id.IsZero() { // Validate provided object id
		return c.JSON(http.StatusBadRequest, interfaces.GenericResponse{
			Error:   true,
			Message: "Provided object id is emtpy or not valid",
		})
	}

	// Get user id from token
	cookie, _ := c.Cookie("access-token")
	token := cookie.Value
	claims := &interfaces.JWTCustomClaims{}
	jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return configs.GetJWTSecret(), nil
	})

	// Verify product exists on cart
	exists, err := models.SearchProductOnCart(claims.ID, payload.Id.Hex())

	if err != nil {
	  return c.JSON(http.StatusInternalServerError, interfaces.GenericResponse{
	    Error: true, 
	    Message: "Unable to find if the product exits",
	  })
	}

	if !exists {
	  return c.JSON(http.StatusNotFound, interfaces.GenericResponse{
	    Error: true, 
	    Message: "Product was nos found on user cart",
	  })
	}

	// Delete from database
	err = models.DeleteCartProduct(claims.ID, payload.Id.Hex())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, interfaces.GenericResponse{
			Error:   true,
			Message: "Product could not be removed from cart.",
		})
	}

	return c.JSON(http.StatusOK, interfaces.GenericResponse{
		Error:   false,
		Message: "Product delete to the cart successfully",
	})
}

// HandleOrderPost creates an order from the items on cart
func HandleOrderPost(c echo.Context) error {
	// Get user id from access token
	cookie, _ := c.Cookie("access-token")
	token := cookie.Value
	claims := &interfaces.JWTCustomClaims{}
	jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return configs.GetJWTSecret(), nil
	})

	// Validate cart length is greater than zero
	cartLength, err := models.GetCartLength(claims.ID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, interfaces.GenericResponse{
			Error:   true,
			Message: "Unable to get user cart from database",
		})
	}

	if cartLength <= 0 {
		return c.JSON(http.StatusNotFound, interfaces.GenericResponse{
			Error:   true,
			Message: "Doesn't found any product on user cart.",
		})
	}

	// Call the stored procedure
	err = models.CreateOrder(claims.ID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, interfaces.GenericResponse{
			Error:   true,
			Message: "Unable to create the order",
		})
	}

	return c.JSON(http.StatusOK, interfaces.GenericResponse{
		Error:   false,
		Message: "Order was created successfully",
	})

}

// obtains the information from cart of a user
func HandleCartGet(c echo.Context) error {
	// Get user id from access token
	cookie, _ := c.Cookie("access-token")
	token := cookie.Value
	claims := &interfaces.JWTCustomClaims{}
	jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return configs.GetJWTSecret(), nil
	})

	var responseCart = []interfaces.CartVerbose{}
	var productInfo = interfaces.CartVerbose{}

	// Verify cart of user
	userCart, err := models.GetCartByUser(claims.ID)
  
	if err != nil {
		return c.JSON(http.StatusInternalServerError, interfaces.GenericResponse{
			Error:   true,
			Message: "Unable to get user cart from database",
		})
	}

	if len(userCart) == 0 {
		return c.JSON(http.StatusNotFound, interfaces.GenericCartProductsResponse{
			Error:    false,
			Message:  "Cart of user empty, nothing to show",
			Products: responseCart,
		})
	}

	//iterating each product
	for _, product := range userCart {
		rows, _ := models.GetDetailsFromID(product.ID.Hex())
		productInfo.Id = rows.ID.Hex()
		productInfo.Name = rows.Name
		productInfo.Units = rows.Units
		productInfo.Quantity = product.Quantity
		productInfo.Price = rows.Price
		productInfo.Image = rows.Image
		responseCart = append(responseCart, productInfo)
	}

	return c.JSON(http.StatusOK, interfaces.GenericCartProductsResponse{
		Error:    false,
		Message:  "Cart of user was generated succesfully",
		Products: responseCart,
	})
}
