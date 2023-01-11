package todo

import (
	"github.com/labstack/echo/v4"
	"go-web-server/models"
	"net/http"
)

func AddTodo(c echo.Context) error {
	req := make(map[string]string)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Can not bind data!")
	}

	todo := models.Todo{
		Name: req["name"],
	}

	if models.CheckExist(todo) {
		return c.JSON(http.StatusBadRequest, "Todo already exist!")
	}

	ok := models.CreateTodo(todo)
	if ok != nil {
		return c.JSON(http.StatusBadRequest, "Can not add todo!")
	}

	return c.JSON(http.StatusOK, "Add todo successfully!")
}
