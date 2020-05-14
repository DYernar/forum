package database

import (
	"fmt"

    _ "github.com/mattn/go-sqlite3"
)


func LikePost(userid int, postid int, like bool) bool {
	db, err := DbConnection()

	if err != nil {
		db.Close()
		fmt.Println(err)
		return false
	}

	if like {
		userliked := GetUserLikes(userid, postid)

		if len(userliked) != 0 && like {
			return true
		}

		userDisliked := GetUserDislikes(userid, postid)
		if len(userDisliked) != 0 {
			_, err := db.Exec("update likes set like=true where userid=$1 and postid=$2", userid, postid)
			if err != nil {
				fmt.Print("liking disliked post err!: ")
				fmt.Println(err)
				return false
			}
			return true
		}

		likes, err := db.Prepare("insert into likes (userid, postid, like) values (?, ?, ?)")

		if err != nil {
			fmt.Print("liking err!: ")
			fmt.Println(err)
			return false
		}
	
		likes.Exec(userid, postid, like)
	} else {
		dislikes := GetUserDislikes(userid, postid)
		if len(dislikes) != 0 {
			return true
		}

		userlikes := GetUserLikes(userid, postid)
		if len(userlikes) != 0 {
			_, err := db.Exec("update likes set like=false where userid=$1 and postid=$2", userid, postid)
			if err != nil {
				fmt.Print("disliking liked post err!: ")
				fmt.Println(err)
				return false
			}
			return true
		}

		likes, err := db.Prepare("insert into likes (userid, postid, like) values (?, ?, ?)")

		if err != nil {
			fmt.Println("like insertion err")
			fmt.Println(err)
			return false
		}
	
		likes.Exec(userid, postid, false)
	}



	db.Close()
	return true
}

func GetUserLikes(postid int, userid int) []int {
	var likes []int
	db, err := DbConnection()
	if err != nil {
		db.Close()
		return likes
	}

	rows, err := db.Query("select userid from likes where userid=$1 and postid=$2 and like=$3", postid, userid, true)

	defer rows.Close()

	if err != nil {
		fmt.Print("like get err!: ")
		fmt.Println(err)
		db.Close()
		return likes
	}

	for rows.Next() {
		var p int
		rows.Scan(&p)
		likes = append(likes, p)
	}
	db.Close()
	return likes
} 



func GetUserDislikes(postid int, userid int) []int {
	var likes []int
	db, err := DbConnection()
	if err != nil {
		db.Close()
		return likes
	}

	rows, err := db.Query("select userid from likes where userid=$1 and postid=$2 and like=$3", postid, userid, false)

	defer rows.Close()

	if err != nil {
		fmt.Print("like get err!: ")
		fmt.Println(err)
		db.Close()
		return likes
	}

	for rows.Next() {
		var p int
		rows.Scan(&p)
		likes = append(likes, p)
	}

	db.Close()
	return likes
} 


func GetLikesByPostID(postid int) []int {
	var likes []int
	db, err := DbConnection()
	if err != nil {
		db.Close()
		return likes
	}

	rows, err := db.Query("select userid from likes where postid=$1 and like=$2", postid, true)

	defer rows.Close()

	if err != nil {
		fmt.Print("like get err!: ")
		fmt.Println(err)
		db.Close()
		return likes
	}

	for rows.Next() {
		var p int
		rows.Scan(&p)
		likes = append(likes, p)
	}

	db.Close()
	return likes
} 

func GetDislikesByPostID(postid int) []int {
	var likes []int
	db, err := DbConnection()
	if err != nil {
		db.Close()
		return likes
	}

	rows, err := db.Query("select userid from likes where postid=$1 and like=$2", postid, false)

	defer rows.Close()

	if err != nil {
		fmt.Print("dislike get err!: ")
		fmt.Println(err)
		db.Close()
		return likes
	}

	for rows.Next() {
		var p int
		rows.Scan(&p)
		likes = append(likes, p)
	}

	db.Close()
	return likes
} 
