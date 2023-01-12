package todo

import (
	"github.com/labstack/echo/v4"
	"go-web-server/models"
	"net/http"
)

func GetListTodo(c echo.Context) error {
	todos, err := models.GetAllTodo()
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Can not get list todos!")
	}

	if len(todos) == 0 {
		return c.JSON(http.StatusOK, "You don't have any todo! Add some more")
	}

	return c.JSONPretty(http.StatusOK, todos, "  ")
}
