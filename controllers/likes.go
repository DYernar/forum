package controller

import(
	"net/http"
	database "forum/database"
	"strconv"
	"fmt"
)


func Like(w http.ResponseWriter, r *http.Request) {
	username, auth := IsAuthorized(r)
		
	if !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	if r.URL.Path == "/like" {
		if r.Method == "GET" {
			if username != "" {
				r.ParseForm()
				postid, _ := strconv.Atoi(r.FormValue("id"))
				user := database.GetUserByName(username)
				database.LikePost(user.UserID, postid, true)
				http.Redirect(w, r, "/#postid"+r.FormValue("id"), http.StatusSeeOther)
			}
		} else {
			BadRequest(w, r)
		}
	} else {
		NotFound(w,r)
	}
}

func Dislike(w http.ResponseWriter, r *http.Request) {
	username, auth := IsAuthorized(r)
		
	if !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	if r.URL.Path == "/dislike" {
		if r.Method == "GET" {
			if username != "" {
				r.ParseForm()
				postid, _ := strconv.Atoi(r.FormValue("id"))
				user := database.GetUserByName(username)
	
				database.LikePost(user.UserID, postid, false)
				http.Redirect(w, r, "/#postid"+r.FormValue("id"), http.StatusSeeOther)
			} else {
				fmt.Println("unauth")
			}
		} else {
			BadRequest(w, r)
		}
	} else {
		NotFound(w,r)
	}
}