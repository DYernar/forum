package controller

import (
	"fmt"
	database "forum/database"
	model "forum/model"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
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
			r.ParseMultipartForm(10 << 20)
			var name string

			file, handler, err := r.FormFile("myFile")
			if err == nil {
				defer file.Close()
				fmt.Print("SIZEL: ")
				if handler.Size > 20127551 {
					var allData model.AllData
					allData.User = database.GetUserByName(uname)
					allData.User.Posts = database.GetPostsByUserID(allData.User.UserID)
					allData.Categories = []string{"Books", "Movies", "Music", "Sport"}
					t := template.Must(template.New("profile").ParseFiles("static/profile.html", "static/header.html"))
					allData.Error = "image size is too big"
					var newPost model.Post
					user := database.GetUserByName(uname)
					newPost.UserID = user.UserID
					newPost.Text = r.FormValue("text")
					newPost.Title = r.FormValue("title")
					fmt.Println(newPost.Text)
					allData.CurrentPost = newPost
					t.Execute(w, allData)
					return
				}

				var extension string
				var arr = strings.Split(handler.Filename, ".")
				extension = arr[len(arr)-1]
				fmt.Print("EXTENSION: ")
				fmt.Println(extension)

				if extension != "jpeg" && extension != "png" && extension != "gif" && extension != "jpg" {
					var allData model.AllData
					allData.User = database.GetUserByName(uname)
					allData.User.Posts = database.GetPostsByUserID(allData.User.UserID)
					allData.Categories = []string{"Books", "Movies", "Music", "Sport"}
					t := template.Must(template.New("profile").ParseFiles("static/profile.html", "static/header.html"))
					allData.Error = "invalid image type it can be only: jpeg, png, gif or jpg"
					var newPost model.Post
					user := database.GetUserByName(uname)
					newPost.UserID = user.UserID
					newPost.Text = r.FormValue("text")
					newPost.Title = r.FormValue("title")
					fmt.Println(newPost.Text)
					allData.CurrentPost = newPost
					t.Execute(w, allData)
					return
				}

				lastPost := database.GetLastPost()
				name = strconv.Itoa(lastPost) + extension

				fileBytes, _ := ioutil.ReadAll(file)

				_ = ioutil.WriteFile(name, fileBytes, 0644)

				_ = os.Rename(name, "./images/"+name)
			}
			r.ParseForm()
			var newPost model.Post
			user := database.GetUserByName(uname)
			newPost.UserID = user.UserID
			newPost.Text = r.FormValue("text")
			newPost.Title = r.FormValue("title")
			newPost.Category = r.FormValue("category")
			newPost.Image = name
			database.InsertPost(newPost)
			http.Redirect(w, r, "/profile", http.StatusSeeOther)

		} else {
			BadRequest(w, r)
		}
	} else {
		NotFound(w, r)
	}
}
