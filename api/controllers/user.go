package controllers

import (
	"net/http"
	"regexp"

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

	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if !re.MatchString(payload.Email) {
		return c.JSON(http.StatusBadRequest, interfaces.GenericResponse{
			Error:   true,
			Message: "This string is not a mail",
		})
	}

	if !models.FindByEmail(payload.Email) {
		return c.JSON(http.StatusBadRequest, interfaces.GenericResponse{
			Error:   true,
			Message: "This mail already exists",
		})
	}

	// Sanitize and save on database
	user := interfaces.Users{
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
	//for test
	//log.Println(user, "user")

	return c.JSON(http.StatusOK, interfaces.GenericResponse{
		Error:   false,
		Message: "User created successfully",
	})
}
