package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
)

type newPostData struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Username string `json:"username"`
	Cookie   string `json:"cookie"`
}

func newpost(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var newpostdata newPostData
	var msg message

	//decoding recieved data
	decoder.Decode(&newpostdata)

	// fmt.Println(newpostdata)

	db, err := sql.Open("sqlite3", "./Database/forum.db")
	checkErr(err)

	check := userCheck(newpostdata.Username, newpostdata.Cookie)

	time := time.Now()

	if check && newpostdata.Title != "" && newpostdata.Content != "" && newpostdata.Category != "" {

		stmt, err := db.Prepare("INSERT INTO post(title, content, created, category,user_name) values(?,?,?,?,?)")
		checkErr(err)

		// adds user to database
		stmt.Exec(newpostdata.Title, newpostdata.Content, time, newpostdata.Category, newpostdata.Username)
		msg.Msg = "Post successfuly created"
		msg.Check = true

	} else {
		msg.Msg = "Please fill all the form"
		msg.Check = false
	}
	defer db.Close()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(msg); err != nil {
		panic(err)
	}
}
