package main

import (
	"database/sql"
)

func updateSession(db *sql.DB, session_id string, username string) {
	var i int

	rows, err := db.Query("SELECT session_id, user_name FROM session WHERE user_name = ?", (username))
	checkErr(err)

	for rows.Next() {
		i++
	}

	if i > 0 {
		stmt, err := db.Prepare("DELETE FROM session WHERE user_name=?")
		checkErr(err)
		stmt.Exec(username)
	}

	stmt, err := db.Prepare(`INSERT INTO session(session_id, user_name, max_age) values(?,?, datetime("now", "+15 minutes"))`)
	checkErr(err)

	stmt.Exec(session_id, username)
}

//Checks if username and sessionId in cookie are the same as in DataBase
func userCheck(username string, sessionId string) bool {
	db, err := sql.Open("sqlite3", "./Database/forum.db")
	checkErr(err)
	if username != "" || sessionId != "" {
		rows, err := db.Query("SELECT session_id, user_name FROM session WHERE user_name = ?", (username))
		checkErr(err)
		user := make([]cookie, 0)

		for rows.Next() { //for loop through database table
			ourUser := cookie{}
			err = rows.Scan(&ourUser.Cookie, &ourUser.Username)
			checkErr(err)
			user = append(user, ourUser)
		}
		defer db.Close()
		if user[0].Cookie == sessionId {
			return true
		}
		return false
	}
	return false
}

//Gets username from database based on sessionID
func getUsername(sessionId string) string {
	db, err := sql.Open("sqlite3", "./Database/forum.db")
	checkErr(err)
	var username string
	if sessionId != "" {
		rows, err := db.Query("SELECT user_name FROM session WHERE session_id = ?", (sessionId))
		checkErr(err)

		for rows.Next() { //for loop through database table
			err = rows.Scan(&username)
			checkErr(err)
			// user = append(user, ourUser)
		}
	}
	defer db.Close()
	return username
}
