package controller

import(
	"net/http"
	database "forum/database"
)


func Drop(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/drop" {
		if r.Method == "GET" {
			database.DropTables()
		} else {
			BadRequest(w, r)
		}
	} else {
		NotFound(w,r)
	}
}