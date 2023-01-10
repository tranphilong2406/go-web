package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Status    int    `json:"status"`
}

var (
	currentTime = time.Now().String()
	timeTrim    = timeConvert(currentTime)
)

func SaveUser(user User) error {
	db, err := sql.Open("mysql", os.Getenv("DB_CONNECTION"))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	query := fmt.Sprintf("insert into `quiz-db`.users(username,email,created_at,updated_at,status) values(\"%s\",\"%s\",\"%s\",\"%s\",%d)", user.Username, user.Email, timeTrim, timeTrim, user.Status)

	insert, err := db.Query(query)
	if err != nil {
		return err
	}

	defer insert.Close()
	return nil
}

func GetUserById(id int) (User, error) {
	db, err := sql.Open("mysql", os.Getenv("DB_CONNECTION"))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var user User

	er := db.QueryRow("SELECT * from `quiz-db`.users where id = ?", id).Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.Status)
	if er != nil {
		panic(er.Error())
	}
	return user, nil
}

func GetUserByEmail(email string) (User, error) {
	db, err := sql.Open("mysql", os.Getenv("DB_CONNECTION"))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var user User

	er := db.QueryRow("SELECT * from `quiz-db`.users where email = ?", email).Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.Status)
	if er != nil {
		return user, er
	}
	return user, nil
}

func UpdateUser(user User) error {
	db, err := sql.Open("mysql", os.Getenv("DB_CONNECTION"))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	_, err = db.Query("update `quiz-db`.users set username = ? , updated_at = ? where id = ?", user.Username, timeTrim, user.ID)

	if err != nil {
		return err
	}

	return nil
}

func timeConvert(s string) string {
	temp := []rune(s)

	newS := temp[:19]
	return string(newS)
}
