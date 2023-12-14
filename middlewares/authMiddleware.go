package middlewares

import (
	"github.com/farhad-aman/cart-web-midterm/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing auth token")
		}

		splitToken := strings.Split(authHeader, "Bearer ")
		if len(splitToken) != 2 {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token format")
		}

		tokenStr := splitToken[1]
		claims, err := utils.ParseToken(tokenStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		c.Set("userID", claims.UserID)
		return next(c)
	}
}
