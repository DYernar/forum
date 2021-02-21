package database

import (
	"fmt"
	model "forum/model"
)

func InsertPost(post model.Post) bool {
	db, err := DbConnection()
	if err != nil {
		db.Close()
		return false
	}

	p, err := db.Prepare("insert into posts(userid, title, text, category, image) values (?, ?, ?, ?, ?)")

	if err != nil {
		fmt.Print("post insertion err!: ")
		fmt.Println(err)
		db.Close()
		return false
	}

	p.Exec(post.UserID, post.Title, post.Text, post.Category, post.Image)

	db.Close()
	return true
}

func GetAllPosts() []model.Post {
	var posts []model.Post
	db, err := DbConnection()
	if err != nil {
		db.Close()
		return posts
	}

	rows, err := db.Query("select rowid, userid, title, text, category, image from posts")

	if err != nil {
		fmt.Print("post insertion err!: ")
		fmt.Println(err)
		db.Close()
		return posts
	}
	defer rows.Close()

	for rows.Next() {
		var p model.Post
		rows.Scan(&p.PostID, &p.UserID, &p.Title, &p.Text, &p.Category, &p.Image)
		p.Comments = GetCommentsByPostID(p.PostID)
		p.Likes = GetLikesByPostID(p.PostID)
		p.Dislikes = GetDislikesByPostID(p.PostID)
		posts = append(posts, p)
	}

	db.Close()
	return posts
}

func GetPostsByUserID(userid int) []model.Post {
	var posts []model.Post
	db, err := DbConnection()
	if err != nil {
		db.Close()
		return posts
	}

	rows, err := db.Query("select rowid, userid, title, text, category, image from posts where userid=$1", userid)

	if err != nil {
		fmt.Print("post insertion err!: ")
		fmt.Println(err)
		db.Close()
		return posts
	}
	defer rows.Close()

	for rows.Next() {
		var p model.Post
		rows.Scan(&p.PostID, &p.UserID, &p.Title, &p.Text, &p.Category, &p.Image)
		p.Comments = GetCommentsByPostID(p.PostID)
		p.Likes = GetLikesByPostID(p.PostID)
		p.Dislikes = GetDislikesByPostID(p.PostID)
		posts = append(posts, p)
	}

	db.Close()
	return posts
}

func GetPostByPostID(postid int) model.Post {
	var p model.Post
	db, err := DbConnection()
	if err != nil {
		db.Close()
		return p
	}

	rows, err := db.Query("select rowid, userid, title, text, category,image from posts where rowid=$1", postid)

	if err != nil {
		fmt.Print("post insertion err!: ")
		fmt.Println(err)
		db.Close()
		return p
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&p.PostID, &p.UserID, &p.Title, &p.Text, &p.Category, &p.Image)
		break
	}

	db.Close()
	return p
}
