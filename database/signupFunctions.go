package database

import (
	"database/sql"
	"fmt"
	model "forum/model"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func UsernameExists(name string) (bool, error) {
	if name == "" {
		return false, nil
	}
	db, err := DbConnection()
	if err != nil {
		fmt.Println("db connection error: ", err.Error())
		return false, err
	}

	err = db.QueryRow("SELECT username FROM users WHERE username = $1", name).Scan(&name)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println("other error")
			defer db.Close()
			return false, err
		}
		fmt.Println(err)
		defer db.Close()
		return false, nil
	}
	defer db.Close()
	return true, nil
}

func EmailExists(email string) (bool, error) {
	if email == "" {
		return false, nil
	}
	db, err := DbConnection()
	if err != nil {
		fmt.Println("db connection error: ", err.Error())
		return false, err
	}

	err = db.QueryRow("SELECT email FROM users WHERE email = $1", email).Scan(&email)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println("other error")
			defer db.Close()
			return false, err
		}
		fmt.Println(err)
		defer db.Close()
		return false, nil
	}
	defer db.Close()
	return true, nil
}

func Signup(user model.User) bool {
	db, err := DbConnection()

	if err != nil {
		return false
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	users, err := db.Prepare("insert into users (username, email, password) values (?, ?, ?)")

	if err != nil {
		fmt.Print("usertable creation err!: ")
		fmt.Println(err)
		return false
	}

	users.Exec(user.Username, user.Email, hashedPassword)

	db.Close()
	return true
}
