package todo

import (
	"github.com/labstack/echo/v4"
	"go-web-server/models"
	"net/http"
	"strconv"
)

func EditTodo(c echo.Context) error {
	req := c.Param("id")
	id, _ := strconv.Atoi(req)
	var myTodo models.Todo

	if err := c.Bind(&myTodo); err != nil {
		return c.JSON(http.StatusBadRequest, "Can not bind data!")
	}

	myTodo.ID = id

	ok := models.EditTodo(myTodo)
	if ok != nil {
		return c.JSON(http.StatusBadRequest, "Can not edit todo!")
	}

	return c.JSON(http.StatusOK, "Update todo successfully!")
}
