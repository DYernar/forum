package controller

import (
	database "forum/database"
	"net/http"
	"strconv"
)

func Comment(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/comment" {
		if r.Method == "POST" {

			username, auth := IsAuthorized(r)

			if !auth {
				http.Redirect(w, r, "/post?id="+r.FormValue("id"), http.StatusSeeOther)
			}

			user := database.GetUserByName(username)
			r.ParseForm()
			postid, _ := strconv.Atoi(r.FormValue("id"))
			comment := r.FormValue("comment")

			database.CommentPost(user.UserID, postid, comment)
			http.Redirect(w, r, "/post?id="+r.FormValue("id"), http.StatusSeeOther)
		} else {
			BadRequest(w, r)
		}
	} else {
		NotFound(w, r)
	}
}
