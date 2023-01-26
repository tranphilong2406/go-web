package models

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
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

func SaveAccount(account Account) error {
	db, err := sql.Open("mysql", os.Getenv("DB_CONNECTION"))
	if err != nil {
		return err
	}

	defer db.Close()

	query := fmt.Sprintf("insert into `quiz-db`.accounts(email,password,user_id,status,created_at,updated_at) values(\"%s\",\"%s\",%d,%d,\"%s\",\"%s\")", account.Email, account.Password, account.UserID, account.Status, time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339))
	insert, e := db.Query(query)
	if e != nil {
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
