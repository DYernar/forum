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
			file, handler, err := r.FormFile("myFile")
			if err != nil {
				fmt.Println("Error Retrieving the File")
				fmt.Println(err)
				return
			}
			defer file.Close()
			fmt.Printf("Uploaded File: %+v\n", handler.Filename)
			fmt.Printf("File Size: %+v\n", handler.Size)
			fmt.Printf("MIME Header: %+v\n", handler.Header)
			var extension string
			for i := 0; i < len(handler.Filename); i++ {
				if handler.Filename[i] == '.' {
					for j := i; j < len(handler.Filename); j++ {
						extension = extension + string(handler.Filename[j])
					}
					break
				}
			}
			if extension != ".jpg" && extension != ".png" && extension != ".gif" && extension != ".jpeg" {
				http.Redirect(w, r, "/profile", http.StatusSeeOther)
				return
			}
			var name string
			lastPost := database.GetLastPost()
			name = strconv.Itoa(lastPost) + extension
			fmt.Print(name)

			fileBytes, err := ioutil.ReadAll(file)

			err = ioutil.WriteFile(name, fileBytes, 0644)
			if err != nil {
				fmt.Println("HEREEEEE")
			}
			err = os.Rename(name, "./images/"+name)
			fmt.Println(err)
			r.ParseForm()
			var newPost model.Post
			user := database.GetUserByName(uname)
			newPost.UserID = user.UserID
			newPost.Text = r.FormValue("text")
			newPost.Title = r.FormValue("title")
			newPost.Category = r.FormValue("category")
			newPost.Image = name
			fmt.Printf(name)
			database.InsertPost(newPost)
			http.Redirect(w, r, "/profile", http.StatusSeeOther)

		} else {
			BadRequest(w, r)
		}
	} else {
		NotFound(w, r)
	}
}
