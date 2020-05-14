package controller

import(
	model "forum/model"
	database "forum/database"
	"net/http"
	"html/template"
	uuid "github.com/satori/go.uuid"
	"time"
)

var cookies map[string]*http.Cookie


func Login(w http.ResponseWriter, r *http.Request) {
	if cookies == nil {
		cookies = map[string]*http.Cookie{}
	}
	
	_, auth := IsAuthorized(r)

	if auth {
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
	}

	if r.URL.Path == "/login" {
		if r.Method == "GET" {
			t := template.Must(template.New("login").ParseFiles("static/login.html", "static/header.html"))
			t.Execute(w, nil)
		} else if r.Method == "POST" {
			r.ParseForm()
			var response model.Response
			response.User.Email = r.FormValue("email")
			password := r.FormValue("password")
			if database.CheckCredentials(response.User.Email, password) {
				response.Success = true
				response.User = database.GetUserByEmail(response.User.Email)
				u:= uuid.NewV4()
				sessionToken := u.String()

				cookie := &http.Cookie{
					Name: "session_token",
					Value: sessionToken, // Some encoded value
					Path: "/", // Otherwise it defaults to the /login if you create this on /login (standard cookie behaviour)
					Expires: time.Now().Add(7200 * time.Second),
				}

				cookies[response.User.Username] = cookie
				http.SetCookie(w, cookie)
				http.Redirect(w, r, "/profile", http.StatusSeeOther)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				response.Success = false
				response.Response = "invalid credentials"
			}
			t := template.Must(template.New("login").ParseFiles("static/login.html", "static/header.html"))
			t.Execute(w, response)
		} else {
			BadRequest(w, r)
		}
	} else {
		NotFound(w,r)
	}
}


func IsAuthorized(r *http.Request) (string, bool) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			return "",false
		}
		return "",false
	}
	for username, cookie := range cookies {
		if cookie.Value == c.Value {
			if cookie.Expires.Sub(time.Now()) <= 0 {

				return "", false
			} else {

				return username, true
			}
		}
	}	
	return "", false
}