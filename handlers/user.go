package handlers

import (
	"github.com/farhad-aman/cart-web-midterm/database"
	"github.com/farhad-aman/cart-web-midterm/models"
	"github.com/farhad-aman/cart-web-midterm/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func RegisterUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	if err := validate.Struct(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := user.HashPassword(user.Password); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to hash password")
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key value") {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Username already exists"})
		}
		return c.JSON(http.StatusInternalServerError, result.Error.Error())
	}

	return c.JSON(http.StatusCreated, user)
}

func LoginUser(c echo.Context) error {
	credentials := new(models.User)
	if err := c.Bind(credentials); err != nil {
		return err
	}

	var user models.User
	database.DB.Where("username = ?", credentials.Username).First(&user)

	if user.ID == 0 || user.CheckPassword(credentials.Password) != nil {
		return c.JSON(http.StatusUnauthorized, "Incorrect username or password")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to generate token")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
