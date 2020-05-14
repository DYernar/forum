package controller

import(
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	uname, auth := IsAuthorized(r)

	if !auth {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	delete(cookies, uname)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}