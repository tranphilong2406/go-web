package main

import (
	"github.com/labstack/echo/v4"
	"go-web-server/api"
	"go-web-server/auth"
	"go-web-server/middlewares"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello world")
	})
	e.POST("/sign-in", auth.SignIn)
	e.POST("/sign-up", auth.SignUp)
	e.POST("/update-user", middlewares.TokenHandle(auth.SignIn))

	e.GET("/get-user/:id", middlewares.TokenHandle(api.GetUser))
	e.Logger.Fatal(e.Start(":8080"))
}
