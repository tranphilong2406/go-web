package auth

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go-web-server/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func SignIn(c echo.Context) error {

	email := c.FormValue("email")
	password := c.FormValue("password")

	myAccount, err := models.GetAccount(email)

	if err != nil {
		return c.JSONPretty(http.StatusBadRequest, "Email not exist!", "  ")
	}

	if !CheckPasswordHash(password, myAccount.Password) {
		return c.JSONPretty(http.StatusOK, "Password does not match!", "  ")
	}

	token, e := CreateJwt(myAccount.UserID)
	if e != nil {
		panic(e)
	}
	res := make(map[string]string)

	res["token"] = token
	return c.JSONPretty(http.StatusOK, res, "  ")
}

func CreateJwt(uid int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(5 * time.Hour)
	claims["uid"] = uid
	key := []byte(os.Getenv("JWT_SECRET_KEY"))
	tokenStr, err := token.SignedString(key)

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
