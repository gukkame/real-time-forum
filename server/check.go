package main

import (
	"database/sql"
	"fmt"
	"log"
)

func getAllPosts(db *sql.DB) []post {
	rows, err := db.Query("SELECT id, title, content, created, category, user_name FROM post")
	checkErr(err)

	postinfo := make([]post, 0)

	for rows.Next() { //for loop through database table
		onePost := post{}
		err = rows.Scan(&onePost.ID, &onePost.Title, &onePost.Content, &onePost.Created, &onePost.Category, &onePost.Username)
		checkErr(err)
		time := ""
		time = onePost.Created[:10]
		time += " " + onePost.Created[11:19]
		onePost.Created = time
		postinfo = append(postinfo, onePost)
	}
	return postinfo
}
func getComments(db *sql.DB, postid int) []comment {
	rows, err := db.Query("SELECT id, content, post_id, user_name FROM comment WHERE post_id = ?", (postid))
	checkErr(err)

	allComments := make([]comment, 0)

	for rows.Next() { //for loop through database table
		onePost := comment{}
		err = rows.Scan(&onePost.ID, &onePost.Content, &onePost.PostId, &onePost.Username)
		checkErr(err)
		allComments = append(allComments, onePost)
	}
	fmt.Println(allComments)
	return allComments
}
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
