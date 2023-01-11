package todo

import (
	"github.com/labstack/echo/v4"
	"go-web-server/models"
	"net/http"
	"strconv"
)

func GetTodo(c echo.Context) error {
	req := c.Param("id")
	id, _ := strconv.Atoi(req)
	myTodo, err := models.GetTodoById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Can not find todo!")
	}

	return c.JSONPretty(http.StatusOK, myTodo, "  ")
}
