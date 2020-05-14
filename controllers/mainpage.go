package controller

import(
	"net/http"
	"html/template"
	model "forum/model"
	database "forum/database"
	"strings"
	"fmt"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		if r.Method == "GET" {
			r.ParseForm()
			title := r.FormValue("title")
			
			var allPosts model.AllData
			allPosts.Posts = database.GetAllPosts()
			allPosts.Categories = []string{"Books", "Movies", "Music", "Sport"}
			for i := 0; i < len(allPosts.Posts); i++ {
				if !strings.HasPrefix(strings.ToLower(allPosts.Posts[i].Title), strings.ToLower(title)) {
					allPosts.Posts = append(allPosts.Posts[:i], allPosts.Posts[i+1:]...)
					i--
				}
			}

			category := r.FormValue("category")

			uname, auth := IsAuthorized(r)

			if auth {
				user := database.GetUserByName(uname)
				likedbyme := r.FormValue("liked")
				myposts := r.FormValue("mypost")

				if likedbyme == "on" {
					for i := 0; i < len(allPosts.Posts); i++ {
						liked := false
						for _, likeid := range allPosts.Posts[i].Likes {
							if likeid == user.UserID {
								liked = true
								break
							}
						} 
						if !liked {
							allPosts.Posts = append(allPosts.Posts[:i], allPosts.Posts[i+1:]...)
							i--
						}
					}
				}

				if myposts == "on" {
					fmt.Println("myposts")
					for i := 0; i < len(allPosts.Posts); i++ {
						if allPosts.Posts[i].UserID != user.UserID {
							allPosts.Posts = append(allPosts.Posts[:i], allPosts.Posts[i+1:]...)
							i--
						}
					}
				}
			}

			if category != "" {
				for i := 0; i < len(allPosts.Posts); i++ {
					if allPosts.Posts[i].Category != category {
						allPosts.Posts = append(allPosts.Posts[:i], allPosts.Posts[i+1:]...)
						i--
					}
				}
			}
			allPosts.LoggedIn = auth			
			t := template.Must(template.New("index").ParseFiles("static/index.html", "static/header.html"))
			t.Execute(w, allPosts)
		} else {
			BadRequest(w, r)
		}
	} else {
		NotFound(w,r)
	}
}