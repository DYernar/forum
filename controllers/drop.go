package controller

import (
	database "forum/database"
	"net/http"
)

func Drop(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/drop" {
		if r.Method == "GET" {
			database.DropTables()
		} else {
			BadRequest(w, r)
		}
	} else {
		NotFound(w, r)
	}
}
