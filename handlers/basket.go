package handlers

import (
	"github.com/farhad-aman/cart-web-midterm/database"
	"github.com/farhad-aman/cart-web-midterm/models"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func validState(fl validator.FieldLevel) bool {
	state := fl.Field().String()
	return state == "PENDING" || state == "COMPLETED"
}

func CreateBasket(c echo.Context) error {
	userID := GetUserIDFromContext(c)
	basketInput := new(models.BasketInput)
	if err := c.Bind(basketInput); err != nil {
		return err
	}

	basket := models.Basket{
		UserID: userID,
		Data:   basketInput.Data,
		State:  basketInput.State,
	}

	if err := validate.Struct(&basket); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result := database.DB.Create(&basket)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}

	return c.JSON(http.StatusCreated, basket)
}

func UpdateBasket(c echo.Context) error {
	userID := GetUserIDFromContext(c)
	basketID, _ := strconv.Atoi(c.Param("id"))

	var existingBasket models.Basket
	if err := database.DB.Where("user_id = ? AND id = ?", userID, basketID).First(&existingBasket).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Basket not found")
	}

	if existingBasket.State == "COMPLETED" {
		return c.JSON(http.StatusForbidden, "Cannot update a completed basket")
	}

	basketInput := new(models.BasketInput)
	if err := c.Bind(basketInput); err != nil {
		return err
	}

	existingBasket.Data = basketInput.Data
	existingBasket.State = basketInput.State

	if err := validate.Struct(&existingBasket); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	database.DB.Save(&existingBasket)
	return c.JSON(http.StatusOK, existingBasket)
}

func GetAllBaskets(c echo.Context) error {
	userID := GetUserIDFromContext(c)
	var baskets []models.Basket
	result := database.DB.Where("user_id = ?", userID).Find(&baskets)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}

	return c.JSON(http.StatusOK, baskets)
}

func GetBasket(c echo.Context) error {
	userID := GetUserIDFromContext(c)
	basketID, _ := strconv.Atoi(c.Param("id"))
	var basket models.Basket
	result := database.DB.Where("user_id = ? AND id = ?", userID, basketID).First(&basket)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, "Basket not found")
	}

	return c.JSON(http.StatusOK, basket)
}

func DeleteBasket(c echo.Context) error {
	userID := GetUserIDFromContext(c)
	basketID, _ := strconv.Atoi(c.Param("id"))
	var basket models.Basket
	if err := database.DB.Where("user_id = ? AND id = ?", userID, basketID).Delete(&basket).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Basket not found")
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Basket deleted"})
}
