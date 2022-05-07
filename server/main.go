package main

import (
	"fmt"
	"net/http"
)

func main() {
	// databaseStart()
	http.HandleFunc("/socket", WsEndpoint)
	http.HandleFunc("/login", login)
	http.HandleFunc("/home/posts", posts)
	http.HandleFunc("/home/singlepost", singlepostcom)
	http.HandleFunc("/home/newcomment", newcomment)
	http.HandleFunc("/newpost", newpost)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/profile", profile)
	fmt.Println("Server is running")
	http.ListenAndServe(":8090", nil)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("JUP ERROR")
		panic(err)
	}
}
