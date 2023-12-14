package main

import (
	"github.com/farhad-aman/cart-web-midterm/database"
	"github.com/farhad-aman/cart-web-midterm/handlers"
	"github.com/farhad-aman/cart-web-midterm/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	handlers.InitValidator()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	database.Connect()

	e.POST("/users/register", handlers.RegisterUser)
	e.POST("/users/login", handlers.LoginUser)

	e.POST("/basket", handlers.CreateBasket, middlewares.JWTMiddleware)
	e.GET("/basket", handlers.GetAllBaskets, middlewares.JWTMiddleware)
	e.GET("/basket/:id", handlers.GetBasket, middlewares.JWTMiddleware)
	e.PATCH("/basket/:id", handlers.UpdateBasket, middlewares.JWTMiddleware)
	e.DELETE("/basket/:id", handlers.DeleteBasket, middlewares.JWTMiddleware)

	e.Logger.Fatal(e.Start(":8080"))
}
