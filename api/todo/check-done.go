package todo

import (
	"github.com/labstack/echo/v4"
	"go-web-server/models"
	"net/http"
	"strconv"
)

func CheckDone(c echo.Context) error {
	req := c.Param("id")
	id, _ := strconv.Atoi(req)

	err := models.CheckIsDone(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Can not check done todo!")
	}

	return c.JSON(http.StatusOK, "Check done successfully!")
}
