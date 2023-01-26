package todo

import (
	"github.com/labstack/echo/v4"
	"go-web-server/models"
	"net/http"
	"strconv"
)

func Create(c echo.Context) error {
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

func CheckDone(c echo.Context) error {
	req := c.Param("id")
	id, ok := strconv.Atoi(req)
	if ok != nil {
		return c.JSON(http.StatusBadRequest, "Type incorrect!")
	}

	err := models.CheckIsDone(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Can not check done todo!")
	}

	return c.JSON(http.StatusOK, "Check done successfully!")
}

func Delete(c echo.Context) error {
	req := c.Param("id")
	id, ok := strconv.Atoi(req)
	if ok != nil {
		return c.JSON(http.StatusBadRequest, "Type incorrect!")
	}

	err := models.DeleteTodo(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Can not delete todo!")
	}
	return c.JSON(http.StatusOK, "Delete todo successfully!")
}

func Edit(c echo.Context) error {
	req := c.Param("id")
	id, e := strconv.Atoi(req)
	if e != nil {
		return c.JSON(http.StatusBadRequest, "Type incorrect!")
	}
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

func GetList(c echo.Context) error {
	todos, err := models.GetAllTodo()
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Can not get list todos!")
	}

	if len(todos) == 0 {
		return c.JSON(http.StatusOK, "You don't have any todo! Add some more")
	}

	return c.JSONPretty(http.StatusOK, todos, "  ")
}

func Get(c echo.Context) error {
	req := c.Param("id")
	id, ok := strconv.Atoi(req)
	if ok != nil {
		return c.JSON(http.StatusBadRequest, "Type incorrect!")
	}
	myTodo, err := models.GetTodoById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Can not find todo!")
	}

	return c.JSONPretty(http.StatusOK, myTodo, "  ")
}
