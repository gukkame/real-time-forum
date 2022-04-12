package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func databaseStart() {

	// os.Remove("./DataBase/forum.db")
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
	createCategoryPost(dbConn)
	createComment(dbConn)
	createReactionPost(dbConn)
	createReactionComment(dbConn)
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
		image blob null,
		user_name varchar(255) not null,
	   FOREIGN KEY (user_name) REFERENCES user_account(user_name)
	 )
	 
		`)
	checkErr(err)
	statement.Exec()

}

func createCategoryPost(dbConn *sql.DB) {
	statement, err := dbConn.Prepare(`
	CREATE TABLE  category_post  (
		post_id integer not null ,
		category_id integer not null,
	   PRIMARY KEY ("post_id", "category_id"),
	   FOREIGN KEY (post_id) REFERENCES post(id),
	   FOREIGN KEY (category_id) REFERENCES category(id)
	 )
	 
		`)
	checkErr(err)
	statement.Exec()
}

func createComment(dbConn *sql.DB) {
	statement, err := dbConn.Prepare(`
	CREATE TABLE comment  (
		id  integer not null PRIMARY KEY AUTOINCREMENT,
		content  text not null,
		image blob null,
	   post_id integer not null,
	   user_name varchar(255) not null,
	   FOREIGN KEY (post_id) REFERENCES post(id),
	   FOREIGN KEY (user_name) REFERENCES user_account(user_name)
	 )
	 `)
	checkErr(err)
	statement.Exec()
}

func createReactionPost(dbConn *sql.DB) {
	statement, err := dbConn.Prepare(`
	CREATE TABLE  reaction_post  (
		emotion integer not null,
	   post_id integer not null,
	   user_name varchar(255) not null,
	 PRIMARY KEY (post_id, user_name)
	   FOREIGN KEY (post_id) REFERENCES post(id),
	   FOREIGN KEY (user_name) REFERENCES user_account(user_name)
	 )
	 
		`)
	checkErr(err)
	statement.Exec()
}

func createReactionComment(dbConn *sql.DB) {
	statement, err := dbConn.Prepare(`
	CREATE TABLE  reaction_comment  (
		emotion integer not null,
		comment_id integer not null,
		user_name varchar(255) not null,
	  PRIMARY KEY (comment_id, user_name)
		FOREIGN KEY (comment_id) REFERENCES comment(id),
		FOREIGN KEY (user_name) REFERENCES user_account(user_name)
	  )
	  
		`)
	checkErr(err)
	statement.Exec()
}
