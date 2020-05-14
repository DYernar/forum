package database

import(
	model "forum/model"
	"fmt"
	"database/sql"
	"golang.org/x/crypto/bcrypt"

)

func GetAllUsers() []model.User {
	db, _ := DbConnection()

	row, err := db.Query("select username, email, password from users")
	fmt.Println(err)
	defer row.Close()
	var all []model.User
	for row.Next() {
		var u model.User
		row.Scan(&u.Username,  &u.Email, &u.Password)
		all = append(all,u)
	}
	db.Close()
	return all
}

func CheckCredentials(email, password string) bool {
	db, err := DbConnection()
	if err != nil {
		fmt.Println("db connection error: ", err.Error())
		return false
	}

	var pass string
	err = db.QueryRow("SELECT password FROM users WHERE email = $1", email).Scan(&pass)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println("other error", err.Error())
			db.Close()
			return false
		}
		db.Close()
		return false
	}
	db.Close()
	return CheckPasswordHash(password, pass)

}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}	

func GetUserByCredentials(username, password string) model.User {
	var user model.User
	db, err := DbConnection()
	if err != nil {
		fmt.Println("db connection error: ", err.Error())
		return user
	}

	row, er := db.Query("SELECT rowid, username, email FROM users WHERE username = $1 and password = $2", username, password)
	if er != nil {
		fmt.Println(er)
	}
	defer row.Close()
	for row.Next() {
		row.Scan(&user.UserID, &user.Username, &user.Email)
	}
	db.Close()
	return user
}


func GetUserByName(username string) model.User {
	var user model.User
	db, err := DbConnection()
	if err != nil {
		db.Close()
		fmt.Println("db connection error: ", err.Error())
		return user
	}

	row, er := db.Query("SELECT rowid, username, email FROM users WHERE username = $1 ", username)
	if er != nil {
		db.Close()
		fmt.Println(er)
	}
	defer row.Close()
	for row.Next() {
		row.Scan(&user.UserID, &user.Username, &user.Email)
	}
	db.Close()
	return user
}

func GetUserByEmail(email string) model.User {
	var user model.User
	db, err := DbConnection()
	if err != nil {
		db.Close()
		fmt.Println("db connection error: ", err.Error())
		return user
	}

	row, er := db.Query("SELECT rowid, username, email FROM users WHERE email = $1 ", email)
	if er != nil {
		db.Close()
		fmt.Println(er)
	}
	defer row.Close()
	for row.Next() {
		row.Scan(&user.UserID, &user.Username, &user.Email)
	}
	db.Close()
	return user
}

func GetUserByID(userid int) model.User {
	var user model.User
	db, err := DbConnection()
	if err != nil {
		db.Close()
		fmt.Println("db connection error: ", err.Error())
		return user
	}

	row, er := db.Query("SELECT rowid, username, email FROM users WHERE rowid = $1 ", userid)
	if er != nil {
		db.Close()
		fmt.Println(er)
	}
	defer row.Close()
	for row.Next() {
		row.Scan(&user.UserID, &user.Username, &user.Email)
	}
	db.Close()
	return user
}