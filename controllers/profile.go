package controller

import(
	"net/http"
	"html/template"
	database "forum/database"
	model "forum/model"
)


func Profile(w http.ResponseWriter, r *http.Request) {

	uname, auth := IsAuthorized(r)

	if !auth {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	if r.URL.Path == "/profile" {
		if r.Method == "GET" {
			var allData model.AllData
			allData.User = database.GetUserByName(uname)
			allData.User.Posts = database.GetPostsByUserID(allData.User.UserID)
			allData.Categories = []string{"Books", "Movies", "Music", "Sport"}
			t := template.Must(template.New("profile").ParseFiles("static/profile.html", "static/header.html"))
			t.Execute(w, allData)
		} else if r.Method == "POST" {
			
			r.ParseForm()
			var newPost model.Post
			user := database.GetUserByName(uname)
			newPost.UserID = user.UserID
			newPost.Text = r.FormValue("text")
			newPost.Title = r.FormValue("title")
			newPost.Category = r.FormValue("category")
			
			database.InsertPost(newPost)
			http.Redirect(w, r, "/profile", http.StatusSeeOther)

		} else {
			BadRequest(w, r)
		}
	} else {
		NotFound(w,r)
	}
}