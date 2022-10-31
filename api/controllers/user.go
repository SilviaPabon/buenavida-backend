package controllers

import (
	"net/http"

	"github.com/SilviaPabon/buenavida-backend/interfaces"
	"github.com/SilviaPabon/buenavida-backend/models"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// HandleUserPost create a new user
func HandleUserPost(c echo.Context) (err error) {
	// Get json palyload
	payload := new(interfaces.Users)

	if err = c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, interfaces.GenericResponse{
			Error:   true,
			Message: "Unable to parse user payload",
		})
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.JSON(http.StatusBadRequest, interfaces.GenericResponse{
			Error:   true,
			Message: "Unable to encrypt the password",
		})
	}

	// Sanitize and save on database
	user := interfaces.Users{
		//id: primitive.NewObjectID(),
		Firstname: payload.Firstname,
		Lastname:  payload.Lastname,
		Email:     payload.Email,
		Password:  string(pass),
	}

	succ := models.SaveUser(&user)

	if !succ {
		return c.JSON(http.StatusInternalServerError, interfaces.GenericResponse{
			Error:   true,
			Message: "Unable to save user on database",
		})
	}

	return c.JSON(http.StatusOK, interfaces.GenericResponse{
		Error:   false,
		Message: "User created successfully",
	})
}
