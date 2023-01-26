package main

import (
	"github.com/labstack/echo/v4"
	"go-web-server/api"
	"go-web-server/api/auth"
	"go-web-server/api/todo"
	"go-web-server/middlewares"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<center><h1>Welcome!</h1></center>")
	})
	e.POST("/sign-in", auth.SignIn)
	e.POST("/sign-up", auth.SignUp)

	e.GET("/get-user/:id", middlewares.TokenHandle(api.GetUser))

	e.POST("/todo", middlewares.TokenHandle(todo.Create))
	e.PUT("/todo/:id", middlewares.TokenHandle(todo.Edit))
	e.GET("/todo", middlewares.TokenHandle(todo.GetList))
	e.GET("/todo/:id", middlewares.TokenHandle(todo.Get))
	e.PUT("/check-done/:id", middlewares.TokenHandle(todo.CheckDone))
	e.DELETE("/todo/:id", middlewares.TokenHandle(todo.Delete))

	e.Logger.Fatal(e.Start(":8080"))
}
