package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type comment struct {
	ID       int    `json:"id"`
	Content  string `json:"content"`
	PostId   int    `json:"postid"`
	Username string `json:"username"`
}

func singlepostcom(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var postid int
	var allComments []comment

	decoder.Decode(&postid)
	fmt.Println("singlepostcomments....", postid)

	db, err := sql.Open("sqlite3", "./Database/forum.db")
	checkErr(err)

	allComments = getComments(db, postid)

	// fmt.Println(allComments)
	defer db.Close()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(allComments); err != nil {
		panic(err)
	}
}
