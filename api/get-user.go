package api

import (
	"github.com/labstack/echo/v4"
	"go-web-server/models"
	"net/http"
	"strconv"
)

func GetUser(c echo.Context) error {
	idRq := c.Param("id")
	id, _ := strconv.Atoi(idRq)

	user, err := models.GetUserById(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, "User not exist!")
	}
	return c.JSONPretty(http.StatusOK, user, "  ")
}
