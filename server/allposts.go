package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type post struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Created  string `json:"postcreated"`
	Category string `json:"category"`
	Username string `json:"username"`
}

func posts(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var postdata []post
	var cookieData cookie

	decoder.Decode(&cookieData)

	db, err := sql.Open("sqlite3", "./Database/forum.db")
	checkErr(err)

	check := userCheck(db, cookieData.Username, cookieData.Cookie)

	fmt.Println("cookie check ", check)

	postdata = getAllPosts(db)

	// fmt.Println(postdata)
	defer db.Close()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(postdata); err != nil {
		panic(err)
	}
}
