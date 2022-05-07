package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func databaseStart() {

	os.Remove("./DataBase/forum.db")
	log.Println("Creating sqlite-database.db...")

	file, err := os.Create("./Database/forum.db") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()

	log.Println("forum.db created")

	dbConn, err := sql.Open("sqlite3", "./Database/forum.db")
	if err != nil {
		// log.Fatal(err)
		checkErr(err)
	}
	createUserAcc(dbConn)
	createSession(dbConn)
	createPost(dbConn)
	createComment(dbConn)
	messages(dbConn)
	defer dbConn.Close() // Defer Closing the database
}

func createUserAcc(dbConn *sql.DB) {
	statement, _ := dbConn.Prepare(`
	CREATE TABLE  user_account  (
		user_name varchar(255) not null PRIMARY KEY,
		age date not null,
		gender varchar(255) not null,
		firstname varchar(255) not null,
		lastname varchar(255) not null,
		email varchar(255) not null,
		password varchar(255) not null
		)
		`)
	statement.Exec()

	stmt, err := dbConn.Prepare(`INSERT INTO user_account(user_name, age, gender, firstname, lastname, email, password) VALUES(?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	stmt.Exec("User1", "2002-09-26", "female", "TestUserName", "TestLastName", "test@test.lv", "t")
}
func createSession(dbConn *sql.DB) {
	statement, _ := dbConn.Prepare(`
	CREATE TABLE  session  (
		session_id  varchar(255) not null PRIMARY KEY,
		user_name varchar(255) not null,
		max_age date not null,
	   FOREIGN KEY (user_name) REFERENCES user_account(user_name)
	 )		
		`)
	statement.Exec()
}

func createPost(dbConn *sql.DB) {
	statement, err := dbConn.Prepare(`
	CREATE TABLE  post  (
		id  INTEGER not null PRIMARY KEY AUTOINCREMENT,
		title varchar(255) not null,
		content  text not null,
		created  datetime not null DEFAULT CURRENT_TIMESTAMP,
		category varchar(255) not null,
		user_name varchar(255) not null,
	   FOREIGN KEY (user_name) REFERENCES user_account(user_name)
	 )
	 
		`)
	checkErr(err)
	statement.Exec()
	// image blob null, add this if you wanna add picture to post
	stmt, err := dbConn.Prepare(`INSERT INTO post(id, title, content, created, category, user_name) VALUES(?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	stmt.Exec("1", "First post on NFT category!", "Test post! How are you! Have a nice day! :)", "2022-05-01 11:30:22", "nft", "Test1User")
}

func createComment(dbConn *sql.DB) {
	statement, err := dbConn.Prepare(`
	CREATE TABLE comment  (
		id  integer not null PRIMARY KEY AUTOINCREMENT,
		content  text not null,
	   post_id integer not null,
	   user_name varchar(255) not null,
	   created  datetime not null DEFAULT CURRENT_TIMESTAMP,
	   FOREIGN KEY (post_id) REFERENCES post(id),
	   FOREIGN KEY (user_name) REFERENCES user_account(user_name)
	 )
	 `)
	checkErr(err)
	statement.Exec()
	stmt, err := dbConn.Prepare(`INSERT INTO comment(id, content, post_id, user_name, created) VALUES(?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	stmt.Exec("1", "First post :)", "1", "Gukka", "2022-05-02 12:30:22")
}
func messages(dbConn *sql.DB) {
	statement, err := dbConn.Prepare(`
	CREATE TABLE messages  (
		id  integer not null PRIMARY KEY AUTOINCREMENT,
	 	content  text not null,
		created  datetime not null DEFAULT CURRENT_TIMESTAMP,
	   	user_1 varchar(255) not null,
	  	user_2 varchar(255) not null
	 )
	 `)
	checkErr(err)
	statement.Exec()
}
