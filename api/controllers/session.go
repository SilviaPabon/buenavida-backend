package controllers

import (
	// "fmt"

	"net/http"
	"time"

	"github.com/SilviaPabon/buenavida-backend/configs"
	"github.com/SilviaPabon/buenavida-backend/interfaces"
	"github.com/SilviaPabon/buenavida-backend/models"
	"github.com/SilviaPabon/buenavida-backend/utils"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// HandleLogin login
func HandleLogin(c echo.Context) error {
	// Get json payload
	var payload = new(interfaces.LoginPayload)
	err := c.Bind(payload)

	if err != nil {
		return c.JSON(http.StatusBadRequest, interfaces.GenericResponse{
			Error:   true,
			Message: "Unable to process query. Try again and make sure mail and password fields were provided",
		})
	}

	// Validate json payload
	if payload.Mail == "" || payload.Password == "" {
		return c.JSON(http.StatusBadRequest, interfaces.GenericResponse{
			Error:   true,
			Message: "Mail and password files are required / can't be empty",
		})
	}

	// Get user from database
	user, err := models.GetUserFromMail(payload.Mail)

	if err != nil {
		return c.JSON(http.StatusNotFound, interfaces.GenericResponse{
			Error:   true,
			Message: "User wasn't found",
		})
	}

	publicUser := interfaces.PublicUser{
		ID:        user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
	}

	// Compare user password
	passwordOk := utils.ComparePasswords([]byte(user.Password), []byte(payload.Password))

	if !passwordOk {
		return c.JSON(http.StatusForbidden, interfaces.GenericResponse{
			Error:   true,
			Message: "Password is not correct",
		})
	}

	// fmt.Printf("%+v\n", user)
	// *** Create tokens ***
	accessToken, _, ATerr := utils.CreateJWTAccessToken(&user)
	refreshToken, refreshTokenUUID, RTerr := utils.CreateJWTRefreshToken(&user)

	if ATerr != nil || RTerr != nil {
		return c.JSON(http.StatusInternalServerError, interfaces.GenericResponse{
			Error:   true,
			Message: "Unable to initialize authentication",
		})
	}

	// *** Save token on redis *** ***
	err = models.SaveRefreshTokenOnRedis(refreshTokenUUID, user.Email)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, interfaces.GenericResponse{
			Error:   true,
			Message: "Unable to cache refresh token",
		})
	}

	// fmt.Println(accessToken)
	// fmt.Println(refreshToken)

	// *** Send tokens on cookies ***
	// *** Prepare cookies ***
	accessCookie := new(http.Cookie)
	accessCookie.Name = "access-token"
	accessCookie.Value = accessToken
	// This should be equal to the one in utils.go file
	accessCookie.Expires = time.Now().Add(15 * time.Minute)
	accessCookie.HttpOnly = true
	accessCookie.Domain = "splendid-piroshki-dade21.netlify.app"
	accessCookie.SameSite = http.SameSiteStrictMode
	accessCookie.Path = "/" // Valid for all paths

	refreshCookie := new(http.Cookie)
	refreshCookie.Name = "refresh-token"
	refreshCookie.Value = refreshToken
	refreshCookie.Expires = time.Now().Add(6 * time.Hour)
	refreshCookie.HttpOnly = true
	refreshCookie.Domain = "splendid-piroshki-dade21.netlify.app"
	accessCookie.SameSite = http.SameSiteStrictMode
	refreshCookie.Path = "/api/session/refresh"

	// *** Send cookies on response ***
	c.SetCookie(accessCookie)
	c.SetCookie(refreshCookie)

	return c.JSON(http.StatusOK, interfaces.LoginResponse{
		Error:   false,
		Message: "User authenticated successfully",
		User:    publicUser,
	})
}

// HandleWhoAmI Get user session information from access token
func HandleWhoAmI(c echo.Context) error {
	// Get token from cookie
	cookie, _ := c.Cookie("access-token")
	token := cookie.Value
	claims := &interfaces.JWTCustomClaims{}

	jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return configs.GetJWTSecret(), nil
	})

	// Query database from user email
	user, err := models.GetUserFromMail(claims.Email)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, interfaces.GenericResponse{
			Error:   true,
			Message: "Unable to get user information.",
		})
	}

	// Get only "public" fields
	puser := interfaces.PublicUser{
		ID:        user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
	}

	return c.JSON(http.StatusOK, interfaces.LoginResponse{
		Error:   false,
		Message: "OK",
		User:    puser,
	})
}

// HandleRefresh creates a new access token if a valid refresh token was provided
func HandleRefresh(c echo.Context) error {

	cookie, _ := c.Cookie("refresh-token")

	// Validate claims
	claims := &interfaces.JWTCustomClaims{}
	signedString := cookie.Value

	jwt.ParseWithClaims(signedString, claims, func(t *jwt.Token) (interface{}, error) {
		return configs.GetJWTSecret(), nil
	})

	accessToken, _, ATerr := utils.CreateJWTAccessToken(&interfaces.User{ID: claims.ID, Email: claims.Email})

	if ATerr != nil {
		return c.JSON(http.StatusInternalServerError, interfaces.GenericResponse{
			Error:   true,
			Message: "Unable to generate accessToken",
		})
	}

	accessCookie := new(http.Cookie)
	accessCookie.Name = "access-token"
	accessCookie.Value = accessToken
	// This should be equal to the one in utils.go file
	accessCookie.Expires = time.Now().Add(15 * time.Minute)
	accessCookie.HttpOnly = true
	accessCookie.Domain = "splendid-piroshki-dade21.netlify.app"
	accessCookie.SameSite = http.SameSiteStrictMode
	accessCookie.Path = "/" // Valid for all paths

	// *** Send cookies on response ***
	c.SetCookie(accessCookie)

	return c.JSON(http.StatusOK, interfaces.GenericResponse{
		Error:   false,
		Message: "OK",
	})

}

// Send expired tokens and remove refresh token from redis
func HandleLogout(c echo.Context) error {
	// *** Get token from cookie
	cookie, _ := c.Cookie("access-token")
	token := cookie.Value
	claims := &interfaces.JWTCustomClaims{}

	jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return configs.GetJWTSecret(), nil
	})

	user := interfaces.User{
		ID:    claims.ID,
		Email: claims.Email,
	}

	// *** Remove refresh token from redis
	err := models.DeleteRefreshTokenFromRedis(user.Email)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, interfaces.GenericResponse{
			Error:   true,
			Message: "Unable to perform redis operation",
		})
	}

	// *** Send tokens as cookies
	accessCookie := new(http.Cookie)
	accessCookie.Name = "access-token"
	accessCookie.Value = "" // Empty string
	accessCookie.Expires = time.Now().Add(2 * time.Second)
	accessCookie.HttpOnly = true
	accessCookie.Domain = "splendid-piroshki-dade21.netlify.app"
	accessCookie.SameSite = http.SameSiteStrictMode
	accessCookie.Path = "/"

	refreshCookie := new(http.Cookie)
	refreshCookie.Name = "refresh-token"
	refreshCookie.Value = "" // Empty string
	refreshCookie.Expires = time.Now().Add(2 * time.Second)
	refreshCookie.HttpOnly = true
	refreshCookie.Domain = "splendid-piroshki-dade21.netlify.app"
	refreshCookie.SameSite = http.SameSiteStrictMode
	refreshCookie.Path = "/api/session/refresh"

	c.SetCookie(accessCookie)
	c.SetCookie(refreshCookie)

	return c.JSON(http.StatusOK, interfaces.GenericResponse{
		Error:   false,
		Message: "Successfully logged out",
	})
}
