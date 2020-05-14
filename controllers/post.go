package controller

import(
	"net/http"
	"html/template"
	database "forum/database"
	model "forum/model"
	"strconv"
)


func Post(w http.ResponseWriter, r *http.Request) {
	var resp model.Result

	_, resp.LoggedIn = IsAuthorized(r)

	if r.URL.Path == "/post" {
		if r.Method == "GET" {
			r.ParseForm()
			postid, _ := strconv.Atoi(r.FormValue("id"))
			resp.Post = database.GetPostByPostID(postid)
			resp.Post.Comments = database.GetCommentsByPostID(postid)
			t := template.Must(template.New("post").ParseFiles("static/post.html", "static/header.html"))
			t.Execute(w, resp)
		} else {
			BadRequest(w, r)
		}
	} else {
		NotFound(w,r)
	}
}