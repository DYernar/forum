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

	rows, err := db.Query("select rowid, userid, comment from comments where postid=$1", postid)

	defer rows.Close()

	if err != nil {
		fmt.Print("comment get err!: ")
		fmt.Println(err)
		db.Close()
		return comments
	}

	for rows.Next() {
		var p model.Comment
		rows.Scan(&p.CommentID, &p.User.UserID, &p.Comment)
		p.User = GetUserByID(p.User.UserID)
		p.Likes = GetCommentLikes(p.CommentID)
		p.Dislikes = GetCommentDislikes(p.CommentID)
		p.PostID = postid
		comments = append(comments, p)
	}

	db.Close()
	return comments
}

func LikeComment(commentid int, like bool, userid int) {
	db, err := DbConnection()

	if err != nil {
		fmt.Println(err)
		db.Close()
		return
	}

	rows, err := db.Query("select like from commentlikes where commentid=$1 and userid=$2", commentid, userid)

	if err != nil {
		fmt.Println(err)
		db.Close()
		return
	}
	defer rows.Close()

	var likes []bool

	for rows.Next() {
		var l bool
		rows.Scan(&l)
		likes = append(likes, l)
	}
	if len(likes) == 0 {

		likes, err := db.Prepare("insert into commentlikes (commentid, userid, like) values (?, ?, ?)")

		if err != nil {
			fmt.Println("commentlike insertion err")
			fmt.Println(err)
			db.Close()
			return
		}

		likes.Exec(commentid, userid, like)
	} else {
		if likes[0] == like {
			db.Close()
			return
		} else {
			_, err := db.Exec("update commentlikes set like=$1 where commentid=$2 and userid=$3", like, commentid, userid)
			if err != nil {
				fmt.Print("updateing the like: ")
				fmt.Println(err)
				db.Close()
				return
			}
			db.Close()
		}
	}
}

func GetCommentLikes(commentid int) int {
	db, err := DbConnection()

	if err != nil {
		fmt.Println(err)
		db.Close()
		return 0
	}

	rows, err := db.Query("select like from commentlikes where commentid=$1 and like=true", commentid)

	if err != nil {
		fmt.Println(err)
		db.Close()
		return 0
	}
	defer rows.Close()

	var likes []bool
	for rows.Next() {
		var l bool
		rows.Scan(&l)
		fmt.Println(l)
		likes = append(likes, l)
	}
	db.Close()
	return len(likes)
}

func GetCommentDislikes(commentid int) int {
	db, err := DbConnection()

	if err != nil {
		fmt.Println(err)
		db.Close()
		return 0
	}

	rows, err := db.Query("select like from commentlikes where commentid=$1 and like=false", commentid)

	if err != nil {
		fmt.Println(err)
		db.Close()
		return 0
	}
	defer rows.Close()

	var likes []bool
	for rows.Next() {
		var l bool
		rows.Scan(&l)
		likes = append(likes, l)
	}
	db.Close()
	return len(likes)
}
