package main

import (
	"fmt"
	controller "forum/controllers"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("go", "vet")
	cmd.Env = append(cmd.Env, "CGO_ENABLED=0")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
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
	http.HandleFunc("/commentlike", controller.CommentLike)
	http.HandleFunc("/drop", controller.Drop)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
