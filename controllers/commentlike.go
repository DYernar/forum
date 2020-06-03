package controller

import (
	"fmt"
	database "forum/database"
	"net/http"
	"strconv"
)

func CommentLike(w http.ResponseWriter, r *http.Request) {

	uname, auth := IsAuthorized(r)

	if !auth {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	if r.URL.Path == "/commentlike" {
		if r.Method == "GET" {
			user := database.GetUserByName(uname)
			r.ParseForm()
			id, err := strconv.Atoi(r.FormValue("id"))
			if err != nil {
				fmt.Println(err)
				BadRequest(w, r)
				return
			}
			val := r.FormValue("like")
			l := false
			if val == "1" {
				l = true //like
			} else if val == "0" {
				l = false //dislike
			} else {
				BadRequest(w, r)
			}

			postid := r.FormValue("postid")
			database.LikeComment(id, l, user.UserID)
			http.Redirect(w, r, "/post?id="+postid, http.StatusSeeOther)
		} else {
			BadRequest(w, r)
		}
	} else {
		NotFound(w, r)
	}
}
