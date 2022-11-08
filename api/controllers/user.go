package controllers

import (
	"net/http"

	"github.com/SilviaPabon/buenavida-backend/configs"
	"github.com/SilviaPabon/buenavida-backend/interfaces"
	"github.com/SilviaPabon/buenavida-backend/models"
	"github.com/SilviaPabon/buenavida-backend/utils"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// HandleUserPost create a new user
func HandleUserPost(c echo.Context) (err error) {
	// Get json payload
	payload := new(interfaces.User)

	v := validator.New()

	if err = c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, interfaces.GenericResponse{
			Error:   true,
			Message: "Unable to parse user payload",
		})
	}

	if c.Bind(payload) == nil {
		return c.JSON(http.StatusBadRequest, interfaces.GenericResponse{
			Error:   true,
			Message: "JSON wasn't provided",
		})
	}

	err_v := v.Struct(payload)
	if err_v != nil {
		return c.JSON(http.StatusBadRequest, interfaces.GenericResponse{
			Error:   true,
			Message: err_v.Error(),
		})
	}

	pass, err := utils.HashPassword([]byte(payload.Password))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, interfaces.GenericResponse{
			Error:   true,
			Message: "Unable to encrypt the password",
		})
	}

	if models.FindByEmail(payload.Email) {
		return c.JSON(http.StatusConflict, interfaces.GenericResponse{
			Error:   true,
			Message: "This mail already exists",
		})
	}

	payload.Password = string(pass)

	succ := models.SaveUser(payload)

	if !succ {
		return c.JSON(http.StatusInternalServerError, interfaces.GenericResponse{
			Error:   true,
			Message: "Unable to save user on database",
		})
	}
	//for test
	//log.Println(user, "user")

	return c.JSON(http.StatusOK, interfaces.GenericResponse{
		Error:   false,
		Message: "User created successfully",
	})
}

func FavoritesPost(c echo.Context) error {
	payload := new(interfaces.ProductIdPayload)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, interfaces.GenericResponse{
			Error:   true,
			Message: "Unable to process request.",
		})
	}

	//Boolean, empty or no
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
	err = models.AddFavorites(claims.ID, payload.Id.Hex())

	if err != nil {
		switch err {
		case interfaces.ErrAlreadyInFavorites:
			return c.JSON(http.StatusConflict, interfaces.GenericResponse{
				Error:   true,
				Message: "Product is already on user favorites.",
			})
		default:
			return c.JSON(http.StatusInternalServerError, interfaces.GenericResponse{
				Error:   true,
				Message: "Unable to be added to favorites. Please try again",
			})
		}
	}

	return c.JSON(http.StatusOK, interfaces.GenericResponse{
		Error:   false,
		Message: "Product added to favorites",
	})

}

// FavoritesGET Get favorires ids
func FavoritesGET(c echo.Context) error {
	// Get user id from token
	cookie, _ := c.Cookie("access-token")
	token := cookie.Value
	claims := &interfaces.JWTCustomClaims{}
	jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return configs.GetJWTSecret(), nil
	})

	favorites, err := models.FavoritesGET(claims.ID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, interfaces.GenericResponse{
			Error:   true,
			Message: "Unable to get user information.",
		})
	}

	return c.JSON(http.StatusOK, interfaces.FavoritesListResponse{
		Error:     false,
		Message:   "OK",
		Favorites: favorites,
	})
}

// HandlelDetailedFavorites Get favorites details
func HandleDetailedFavorites(c echo.Context) error {
  // *** Get user data from token
  cookie, _ := c.Cookie("access-token")
  token := cookie.Value
  claims := &interfaces.JWTCustomClaims{}
  jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
    return configs.GetJWTSecret(), nil
  })

  // *** Query database
  favorites, err := models.GetDetailedFavorites(claims.ID)

  if err != nil {
    return c.JSON(http.StatusInternalServerError, interfaces.GenericResponse{
      Error: true, 
      Message: "Unable to get favorites list",
    })
  }

  return c.JSON(http.StatusOK, interfaces.FavoritesDetailsResponse{
    Error: false, 
    Message: "OK",
    Favorites: favorites,
  })
}

// HandleGetOrders Get user orders resume
func HandleGetOrders(c echo.Context) error {
  // *** Get user data from token
  cookie, _ := c.Cookie("access-token")
  token := cookie.Value
  claims := &interfaces.JWTCustomClaims{}
  jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
    return configs.GetJWTSecret(), nil
  })

  // *** Query database
  orders, err := models.GetUserOrdersResume(claims.ID)

  if err != nil {
    return c.JSON(http.StatusInternalServerError, interfaces.GenericResponse{
      Error: true, 
      Message: "Unable to get orders from database",
    })
  }

  return c.JSON(http.StatusOK, interfaces.OrdersResumeResponse{
    Error: false, 
    Message: "Ok",
    Orders: orders,
  })

}
