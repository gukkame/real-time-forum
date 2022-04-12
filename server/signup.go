package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type userData struct {
	Username  string `json:"username"`
	Age       string `json:"age"`
	Gender    string `json:"gender"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type message struct {
	Msg   string `json:"msg"`
	Check bool   `json:"check"`
}

func signup(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var sigupData userData
	var msg message

	//decoding recieved data
	decoder.Decode(&sigupData)

	// fmt.Println(sigupData)

	db, err := sql.Open("sqlite3", "./Database/forum.db")
	checkErr(err)

	//search at database if user already exists with this username or email
	people := searchForUser(db, sigupData.Username, sigupData.Email)
	fmt.Printf("Found %v results\n", len(people))

	if len(people) == 0 {
		hashedpass, err := bcrypt.GenerateFromPassword([]byte(sigupData.Password), bcrypt.MinCost)
		checkErr(err)

		stmt, err := db.Prepare("INSERT INTO user_account(user_name, age, gender,firstname, lastname, email, password) values(?,?,?,?,?,?,?)")
		checkErr(err)

		// adds user to database
		stmt.Exec(sigupData.Username, sigupData.Age, sigupData.Gender, sigupData.Firstname, sigupData.Lastname, sigupData.Email, hashedpass)
		msg.Msg = "User successfuly created"
		msg.Check = true
	} else {
		msg.Msg = "User with this email or username already exist"
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
