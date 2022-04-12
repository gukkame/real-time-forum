package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type loginDataRes struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Time     string `json:"date"`
}

type loginDataSend struct {
	Username string `json:"username"`
	Msg      string `json:"msg"`
	Check    bool   `json:"check"`
	Value    string `json:"value"`
	MaxAge   int    `json:"maxage"`
}

func login(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var loginDataRes loginDataRes
	var loginDataSend loginDataSend

	//decoded recieved data
	decoder.Decode(&loginDataRes)
	fmt.Println("data from form LOGIN", loginDataRes.Username)

	loginDataSend.Msg = ""
	loginDataSend.Username = loginDataRes.Username
	db, err := sql.Open("sqlite3", "./Database/forum.db")
	checkErr(err)

	people := searchForUser(db, loginDataRes.Username, loginDataRes.Username)
	fmt.Printf("Found %v results\n", len(people))
	if len(people) == 0 {
		loginDataSend.Msg = "This user does not exist"
	} else {
		errComparePass := bcrypt.CompareHashAndPassword(people[0].Password, []byte(loginDataRes.Password))
		if errComparePass != nil || loginDataRes.Password == "" {
			loginDataSend.Msg = "Wrong password or username"
			loginDataSend.Check = false
		} else {
			loginDataSend.Msg = "Correct"
			loginDataSend.Check = true
			sID := uuid.NewV4()
			loginDataSend.Value = sID.String()
			loginDataSend.MaxAge = 3 * 60

			// adds user to database
			updateSession(db, loginDataSend.Value, loginDataRes.Username, loginDataRes.Time)

		}
	}
	defer db.Close()

	fmt.Println("data i am sending", loginDataRes)
	fmt.Println()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(loginDataSend); err != nil {
		panic(err)
	}

}
