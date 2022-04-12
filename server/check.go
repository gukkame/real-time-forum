package main

import (
	"database/sql"
	"log"
)

func getUserInfo(db *sql.DB, username string) []userData {
	rows, err := db.Query("SELECT user_name,age,gender,firstname,lastname, email FROM user_account WHERE user_name = ?", (username))
	checkErr(err)

	userinfo := make([]userData, 0)

	for rows.Next() { //for loop through database table
		ourPerson := userData{}
		err = rows.Scan(&ourPerson.Username, &ourPerson.Age, &ourPerson.Gender, &ourPerson.Firstname, &ourPerson.Lastname, &ourPerson.Email)
		checkErr(err)
		userinfo = append(userinfo, ourPerson)
	}
	return userinfo
}

func searchForUser(db *sql.DB, searchUsername string, searchEmail string) []user {

	rows, err := db.Query("SELECT user_name, password, email FROM user_account WHERE user_name = ?", (searchUsername))
	checkErr(err)

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	people := make([]user, 0)

	for rows.Next() { //for loop through database table
		ourPerson := user{}
		err = rows.Scan(&ourPerson.UserName, &ourPerson.Password, &ourPerson.Email)
		checkErr(err)
		people = append(people, ourPerson)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	rows1, err := db.Query("SELECT user_name, password, email FROM user_account WHERE email = ?", (searchEmail))
	checkErr(err)

	err = rows1.Err()
	if err != nil {
		log.Fatal(err)
	}

	for rows1.Next() { //for loop through database table
		ourPerson := user{}
		err = rows1.Scan(&ourPerson.UserName, &ourPerson.Password, &ourPerson.Email)
		checkErr(err)
		people = append(people, ourPerson)
	}

	err = rows1.Err()
	if err != nil {
		log.Fatal(err)
	}

	return people
}
