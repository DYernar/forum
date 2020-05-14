package controller

import(
	"net/http"
	"html/template"
	database "forum/database"
	model "forum/model"
	"fmt"
)


func Signup(w http.ResponseWriter, r *http.Request) {

	_, auth := IsAuthorized(r)

	if auth {
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
	}

	if r.URL.Path == "/signup" {
		if r.Method == "GET" {
			t := template.Must(template.New("signup").ParseFiles("static/signup.html", "static/header.html"))
			t.Execute(w, nil)
		} else if r.Method == "POST" {
			r.ParseForm()
			var response model.Response
			var newUser model.User
			newUser.Username = r.FormValue("username")
			newUser.Password = r.FormValue("password")
			newUser.Email = r.FormValue("email")
			name, err := database.UsernameExists(newUser.Username)
			email, err := database.EmailExists(newUser.Email)
			if !name && !email && err == nil{
				if database.Signup(newUser) {
					fmt.Println("registered")
					response.Response = "Sucess"
					response.User = newUser
					response.Success = true
				} else {
					fmt.Println("not registered")
					response.Response = "Not Sucess"
					response.User = newUser
					response.Success = false
				}
			} else {
				fmt.Println("invalid username email")
				response.EmailExists = email
				response.UsernameExists = name
				response.Error = err
				response.Success = false
			}
			response.User = newUser
			t := template.Must(template.New("signup").ParseFiles("static/signup.html", "static/header.html"))
			t.Execute(w, response)
		} else {
			BadRequest(w, r)
		}
	} else {
		NotFound(w,r)
	}
}


