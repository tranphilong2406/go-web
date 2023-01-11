package todo

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go-web-server/models"
	"net/http"
	"strconv"
)

func EditTodo(c echo.Context) error {
	req := make(map[string]string)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Can not bind data!")
	}

	id, _ := strconv.Atoi(req["id"])

	todo := models.Todo{
		ID:   id,
		Name: req["name"],
	}
	fmt.Println(todo)
	return nil
}
