package models

import (
	"database/sql"
	"os"
	"time"
)

type Todo struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	IsDone    int    `json:"is-done"`
	CreatedAt string `json:"created-at"`
	UpdatedAt string `json:"updated-at"`
}

func CreateTodo(todo Todo) error {
	db, err := sql.Open("mysql", os.Getenv("DB_CONNECTION"))
	if err != nil {
		return err
	}
	defer db.Close()
	currentTime := time.Now().String()
	timeTrim := timeConvert(currentTime)

	insert, ok := db.Query("INSERT INTO `quiz-db`.`to-do`(name, isdone, created_at,updated_at) values (?,?,?,?)", todo.Name, 0, timeTrim, timeTrim)
	if ok != nil {
		return ok
	}
	defer insert.Close()
	return nil
}

func CheckExist(todo Todo) bool {
	db, err := sql.Open("mysql", os.Getenv("DB_CONNECTION"))
	if err != nil {
		return false
	}

	defer db.Close()

	var myTodo Todo

	err = db.QueryRow("SELECT * FROM `quiz-db`.`to-do` WHERE name = ?", todo.Name).Scan(&myTodo.ID, &myTodo.Name, &myTodo.IsDone, &myTodo.CreatedAt, &myTodo.UpdatedAt)
	if err != nil {
		return false
	}
	return true
}

func GetAllTodo() ([]Todo, error) {
	db, err := sql.Open("mysql", os.Getenv("DB_CONNECTION"))
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var todos []Todo

	result, ok := db.Query("SELECT * FROM `quiz-db`.`to-do` WHERE isdone = 1")
	if ok != nil {
		return nil, ok
	}

	for result.Next() {
		var todo Todo
		er := result.Scan(&todo.ID, &todo.Name, &todo.IsDone, &todo.CreatedAt, &todo.UpdatedAt)
		if er != nil {
			return nil, er
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func GetTodoById(id int) (Todo, error) {
	var todo Todo
	db, err := sql.Open("mysql", os.Getenv("DB_CONNECTION"))
	if err != nil {
		return todo, err
	}
	defer db.Close()
	err = db.QueryRow("SELECT * FROM `quiz-db`.`to-do` WHERE id = ?", id).Scan(&todo.ID, &todo.Name, &todo.IsDone, &todo.CreatedAt, &todo.UpdatedAt)

	if err != nil {
		return todo, err
	}

	return todo, nil
}

func CheckIsDone(id int) error {
	db, err := sql.Open("mysql", os.Getenv("DB_CONNECTION"))
	if err != nil {
		return err
	}
	defer db.Close()
	currentTime := time.Now().String()
	timeTrim := timeConvert(currentTime)

	_, ok := db.Query("UPDATE `quiz-db`.`to-do` SET isdone = 1, updated_at = ? WHERE id = ?", timeTrim, id)
	if ok != nil {
		return ok
	}

	return nil
}

func DeleteTodo(id int) error {
	db, err := sql.Open("mysql", os.Getenv("DB_CONNECTION"))
	if err != nil {
		return err
	}
	defer db.Close()

	_, ok := db.Query("DELETE FROM `quiz-db`.`to-do` WHERE id = ?", id)
	if ok != nil {
		return ok
	}
	return nil
}

func EditTodo(todo Todo) error {
	db, err := sql.Open("mysql", os.Getenv("DB_CONNECTION"))
	if err != nil {
		return err
	}
	defer db.Close()

	currentTime := time.Now().String()
	timeTrim := timeConvert(currentTime)

	_, ok := db.Query("UPDATE `quiz-db`.`to-do`SET name = ?, updated_at = ? WHERE id = ?", todo.Name, timeTrim, todo.ID)
	if ok != nil {
		return ok
	}
	return nil
}
