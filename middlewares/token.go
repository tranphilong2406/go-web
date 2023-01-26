package middlewares

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"os"
	"strings"
)

func TokenHandle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request().Header.Get("Authorization")

		if !strings.HasPrefix(req, "Bearer") {
			return c.JSON(401, "Unauthorized!")
		}

		tokenStr := strings.Split(req, "Bearer ")

		token, _ := jwt.Parse(tokenStr[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing")
			}
			sampleSecretKey := os.Getenv("JWT_SECRET_KEY")
			return sampleSecretKey, nil
		})
		claims := token.Claims.(jwt.MapClaims)
		c.Set("uid", claims["uid"])
		return next(c)
	}
}
