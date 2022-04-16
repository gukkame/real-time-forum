package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type newComData struct {
	Content  string `json:"comContent"`
	PostId   int    `json:"postid"`
	Username string `json:"username"`
	Cookie   string `json:"cookie"`
}

func newcomment(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var newcomdata newComData
	var msg message

	//decoding recieved data
	decoder.Decode(&newcomdata)

	fmt.Println("newcoment data  ", newcomdata)

	db, err := sql.Open("sqlite3", "./Database/forum.db")
	checkErr(err)

	check := userCheck(db, newcomdata.Username, newcomdata.Cookie)

	fmt.Println(check)

	if check && newcomdata.Content != "" && newcomdata.Username != "" && newcomdata.PostId != 0 {

		stmt, err := db.Prepare("INSERT INTO comment(content,  post_id, user_name) values(?,?,?)")
		checkErr(err)

		// adds user to database
		stmt.Exec(newcomdata.Content, newcomdata.PostId, newcomdata.Username)
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
