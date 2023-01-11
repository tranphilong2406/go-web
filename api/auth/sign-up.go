package auth

import (
	"github.com/labstack/echo/v4"
	"go-web-server/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/mail"
	"strings"
)

func SignUp(c echo.Context) error {
	req := make(map[string]string)

	if isOk := c.Bind(&req); isOk != nil {
		return isOk
	}

	email := req["email"]
	username := req["username"]
	password := req["password"]
	cfPass := req["cfPass"]

	if !valid(email) {
		return c.JSON(http.StatusBadRequest, "Invalid email")
	}

	if strings.Compare(password, cfPass) != 0 {
		return c.JSON(http.StatusBadRequest, "Password and Confirm password does not match!")
	}

	_, e := models.GetUserByEmail(email)
	if e == nil {
		return c.JSON(http.StatusBadRequest, "Email existed!")
	}

	user := models.User{
		Email:    email,
		Username: username,
		Status:   1,
	}

	err := models.SaveUser(user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	currentUser, er := models.GetUserByEmail(user.Email)
	if er != nil {
		return c.JSON(http.StatusBadRequest, er)
	}

	hashPass, _ := HashPassword(password)

	account := models.Account{
		Email:    email,
		Password: hashPass,
		UserID:   currentUser.ID,
		Status:   1,
	}

	saveAcc := models.SaveAccount(account)

	if saveAcc != nil {
		return c.JSON(http.StatusBadRequest, saveAcc)
	}

	token, ok := CreateJwt(currentUser.ID)
	if ok != nil {
		return c.JSON(http.StatusBadRequest, e)
	}

	res := make(map[string]string)
	res["token"] = token

	return c.JSONPretty(http.StatusCreated, res, "  ")
}

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
