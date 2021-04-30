package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func DbConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Print("DB connection err!: ")
		fmt.Println(err)
		db.Close()
		return db, err
	}

	users, err := db.Prepare("CREATE TABLE IF NOT EXISTS users(username varchar(255), email varchar, password varchar)")

	if err != nil {
		fmt.Print("user table creation err!: ")
		fmt.Println(err)
		db.Close()
		return db, err
	}

	users.Exec()

	posts, err := db.Prepare("CREATE TABLE IF NOT EXISTS posts(userid int, title varchar, text varchar, category varchar,image varchar)")

	if err != nil {
		fmt.Print("post table creation err!: ")
		fmt.Println(err)
		db.Close()
		return db, err
	}

	posts.Exec()

	comments, err := db.Prepare("CREATE TABLE IF NOT EXISTS comments(userid int, postid int, comment varchar)")

	if err != nil {
		fmt.Print("comments table creation err!: ")
		fmt.Println(err)
		db.Close()
		return db, err
	}

	comments.Exec()

	likes, err := db.Prepare("CREATE TABLE IF NOT EXISTS likes(userid int, postid int, like bool)")

	if err != nil {
		fmt.Print("likes table creation err!: ")
		fmt.Println(err)
		db.Close()
		return db, err
	}

	likes.Exec()

	commentLikes, err := db.Prepare("CREATE TABLE IF NOT EXISTS commentlikes(commentid int, userid int, like bool)")

	if err != nil {
		fmt.Print("likes table creation err!: ")
		fmt.Println(err)
		db.Close()
		return db, err
	}

	commentLikes.Exec()

	return db, nil
}

func DropTables() {
	db, _ := DbConnection()

	_, _ = db.Exec("drop table users")
	_, _ = db.Exec("drop table posts")
	_, _ = db.Exec("drop table comments")
	_, _ = db.Exec("drop table likes")
	_, _ = db.Exec("drop table commentlikes")

	db.Close()
}
