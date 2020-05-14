package main

import(
	"net/http"
	"os"
	"forum/controllers"
	"fmt"
)


func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "7070"
	}

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.HandleFunc("/", controller.MainPage)
	http.HandleFunc("/profile", controller.Profile)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/signup", controller.Signup)
	http.HandleFunc("/logout", controller.Logout)
	http.HandleFunc("/post", controller.Post)
	http.HandleFunc("/comment", controller.Comment)
	http.HandleFunc("/like", controller.Like)
	http.HandleFunc("/dislike", controller.Dislike)
	http.HandleFunc("/drop", controller.Drop)


	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
	}
}