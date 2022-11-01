package controllers

import (
	"net/http"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/SilviaPabon/buenavida-backend/interfaces"
	"github.com/SilviaPabon/buenavida-backend/models"
	"github.com/SilviaPabon/buenavida-backend/utils"
)

// HandleUserPost create a new user
func HandleUserPost(c echo.Context) (err error) {
	// Get json payload
	payload := new(interfaces.Users)

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
		return c.JSON(http.StatusBadRequest, interfaces.GenericResponse{
			Error:   true,
			Message: "Unable to encrypt the password",
		})
	}

	if models.FindByEmail(payload.Email) {
		return c.JSON(http.StatusBadRequest, interfaces.GenericResponse{
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
