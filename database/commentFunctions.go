package database

import (
	"fmt"
	model "forum/model"

    _ "github.com/mattn/go-sqlite3"
)


func CommentPost(userid int, postid int, comment string) bool {
	db, err := DbConnection()

	if err != nil {
		db.Close()
		fmt.Println(err)
		return false
	}

	users, err := db.Prepare("insert into comments (userid, postid, comment) values (?, ?, ?)")

	if err != nil {
		fmt.Print("usertable creation err!: ")
		fmt.Println(err)
		return false
	}

	fmt.Println(users.Exec(userid, postid, comment))

	db.Close()
	return true
}


func GetCommentsByPostID(postid int) []model.Comment {
	var comments []model.Comment
	db, err := DbConnection()
	if err != nil {
		db.Close()
		return comments
	}

	rows, err := db.Query("select userid, comment from comments where postid=$1", postid)

	defer rows.Close()

	if err != nil {
		fmt.Print("comment get err!: ")
		fmt.Println(err)
		db.Close()
		return comments
	}

	for rows.Next() {
		var p model.Comment
		rows.Scan(&p.User.UserID, &p.Comment)
		p.User = GetUserByID(p.User.UserID)
		comments = append(comments, p)
	}

	db.Close()
	return comments
} 
