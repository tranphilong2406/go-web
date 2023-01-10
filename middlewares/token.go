package middlewares

import (
	"github.com/labstack/echo/v4"
	"strings"
)

func TokenHandle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")

		if !strings.HasPrefix(token, "Bearer") {
			return c.JSON(401, "Unauthorization!")
		}

		return next(c)
	}
}
