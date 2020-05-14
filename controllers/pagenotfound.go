package controller

import(
	"net/http"
	"html/template"
)


func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	t := template.Must(template.New("404").ParseFiles("static/404.html", "static/header.html"))
	t.Execute(w, nil)
}


func BadRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
	t := template.Must(template.New("500").ParseFiles("static/500.html", "static/header.html"))
	t.Execute(w, nil)
}