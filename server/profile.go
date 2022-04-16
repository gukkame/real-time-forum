package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type cookie struct {
	Cookie   string `json:"cookie"`
	Username string `json:"username"`
}

func profile(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var userdata userData
	var cookieData cookie

	decoder.Decode(&cookieData)

	db, err := sql.Open("sqlite3", "./Database/forum.db")
	checkErr(err)

	check := userCheck(db, cookieData.Username, cookieData.Cookie)

	if check {
		//search at database if user already exists with this username or email
		people := getUserInfo(db, cookieData.Username)
		userdata = userData{
			Username:  people[0].Username,
			Age:       people[0].Age,
			Gender:    people[0].Gender,
			Firstname: people[0].Firstname,
			Lastname:  people[0].Lastname,
			Email:     people[0].Email,
		}

	}
	defer db.Close()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(userdata); err != nil {
		panic(err)
	}
}
