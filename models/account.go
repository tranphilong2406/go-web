package models

import (
	"database/sql"
	"fmt"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

type Account struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	UserID    int    `json:"uid"`
	Status    int    `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CustomValidator struct {
	validator *validator.Validate
}

func SaveAccount(account Account) error {
	db, err := sql.Open("mysql", os.Getenv("DB_CONNECTION"))
	if err != nil {
		return err
	}

	defer db.Close()

	currentTime := time.Now().String()

	timeTrim := timeConvert(currentTime)
	query := fmt.Sprintf("insert into `quiz-db`.accounts(email,password,user_id,status,created_at,updated_at) values(\"%s\",\"%s\",%d,%d,\"%s\",\"%s\")", account.Email, account.Password, account.UserID, account.Status, timeTrim, timeTrim)

	insert, err := db.Query(query)
	if err != nil {
		return err
	}

	defer insert.Close()
	return nil
}

func GetAccount(email string) (Account, error) {
	db, err := sql.Open("mysql", os.Getenv("DB_CONNECTION"))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var account Account

	er := db.QueryRow("SELECT * from `quiz-db`.accounts where email = ?", email).Scan(&account.ID, &account.Email, &account.Password, &account.UserID, &account.Status, &account.CreatedAt, &account.UpdatedAt)
	if er != nil {
		return account, er
	}
	return account, nil
}
