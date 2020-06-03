package model

import (
	"time"
)

type User struct {
	UserID   int
	Username string
	Email    string
	Password string
	Posts    []Post
}

type Response struct {
	Success        bool
	Response       string
	User           User
	EmailExists    bool
	UsernameExists bool
	Error          error
}

type Post struct {
	PostID   int
	UserID   int
	Title    string
	Text     string
	Category string
	Added    time.Time
	Comments []Comment
	Likes    []int
	Dislikes []int
}

type AllData struct {
	LoggedIn   bool
	User       User
	Posts      []Post
	Categories []string
}

type Result struct {
	LoggedIn bool
	Post     Post
}

type Comment struct {
	CommentID int
	PostID    int
	User      User
	Comment   string
	Likes     int
	Dislikes  int
}
