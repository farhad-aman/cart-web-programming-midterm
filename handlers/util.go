package handlers

import (
	"github.com/labstack/echo/v4"
)

func GetUserIDFromContext(c echo.Context) uint {
	userID, ok := c.Get("userID").(uint)
	if !ok {
		panic("Failed to get user ID from context")
	}
	return userID
}
