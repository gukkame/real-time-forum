package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

type RecievedData struct {
	Content   string `json:"content"`
	SessionId string `json:"sessionID"`
	User1     string `json:"user1"`
	User2     string `json:"user2"`
}
type SendMsg struct {
	Content  string `json:"content"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Created  string `json:"created"`
	Date     string `json:"date"`
	Time     string `json:"time"`
}

type SendNotification struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
}

type UsersOnline struct {
	Username string
	Conn     *websocket.Conn
	LastMsg  string
	NoMsg    string
}
type AllUsersOnline struct {
	Allusers []UsersOnline `json:"allUsers"`
}

var (
	wsUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	allUsersOnline []UsersOnline
)

func WsEndpoint(w http.ResponseWriter, r *http.Request) {

	var username string
	var activechatUsers []string
	reqHeader := r.Header.Get("Cookie")
	if reqHeader != "" {
		splitToken := strings.Split(reqHeader, "session=")
		session := strings.Split(splitToken[1], ":")
		username = session[1]
	}
	wsConn, err := Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	useronline := UsersOnline{Username: username, Conn: wsConn}

	// event loop
	if reqHeader != "" {
		check := false
		for _, user := range allUsersOnline {
			if username == user.Username {
				check = true
			}
		}
		if !check {
			allUsersOnline = append(allUsersOnline, useronline)
		}
		// SendAllUsers(AllUsersOnline{allUsersOnline}, wsConn)
	}
	go func() {
		for {
			// var sendMsgUsers AllUsersOnline

			var msg RecievedData

			err := wsConn.ReadJSON(&msg)
			fmt.Println("Recieved Data: ", msg)

			//Get Message history from database
			if msg.Content == "getUserList" {
				lastMsg := getLastMsg(username, allUsersOnline)

				var user UsersOnline
				var users []UsersOnline
				count := 0
				for i, user2 := range allUsersOnline {
					if username != user2.Username {

						user.Username = user2.Username
						user.Conn = user2.Conn
						if lastMsg[i-count] == "0" {
							user.NoMsg = "0"
							user.LastMsg = ""
						} else {

							user.LastMsg = lastMsg[i-count]
							user.NoMsg = ""
						}
						users = append(users, user)
					} else {
						count++
					}
				}
				sendAllUsers(AllUsersOnline{Allusers: users}, wsConn)
			}
			if msg.Content == "chatCloses" {
				for _, user := range activechatUsers {
					if user == msg.User2 {
						// activechatUsers = append(activechatUsers[:i], activechatUsers[i+1:]...)
						activechatUsers = []string{}
					}
				}
			}
			if msg.Content == "check" {
				check := userCheck(msg.User1, msg.SessionId)
				fmt.Println("User check!", check)
				if !check {
					sendCheck(message{Msg: "SessionIDs does not match", Check: false}, wsConn)
				} else {
					sendCheck(message{Msg: "", Check: true}, wsConn)
				}
			}

			//Save message in the database
			if msg.Content != "" && msg.SessionId != "" && msg.User1 != "" && msg.User2 != "" {
				var conn2 *websocket.Conn
				for _, user := range allUsersOnline {
					fmt.Println("Save Message! ", user)
					if user.Username == msg.User2 {
						conn2 = user.Conn
					}
				}
				saveMessage(msg, conn2)
				sendMessage(msg, wsConn, conn2)
				fmt.Println("Save Message! ")

			}
			if msg.Content == "" && msg.SessionId == "" && msg.User1 == "" && msg.User2 != "" {
				activechatUsers = append(activechatUsers, msg.User2)
				sendMessageHistory(username, wsConn, activechatUsers)
				fmt.Println("Send message history!")
				fmt.Println("Active Chats\n ", activechatUsers)
				fmt.Println()
			}

			if err != nil {
				for i, userOnline := range allUsersOnline {
					if userOnline.Username == useronline.Username {
						allUsersOnline = append(allUsersOnline[:i], allUsersOnline[i+1:]...)
					}
				}
				fmt.Printf("error reading JSON: %s\n", err.Error())
				break
			}
		}

	}()

}
func getLastMsg(user1 string, allUsersOn []UsersOnline) []string {
	fmt.Println("User : ", user1)

	db, err := sql.Open("sqlite3", "./Database/forum.db")
	checkErr(err)
	var lastMsg []string

	for _, user2 := range allUsersOnline {
		if user1 != user2.Username {
			var allMsg []string
			fmt.Println("User2(receiver) : ", user2.Username)
			rows, err := db.Query("SELECT created FROM messages WHERE user_1 = ? AND user_2 = ? OR user_1 = ? AND user_2 = ?", user1, user2.Username, user2.Username, user1)
			if err != nil {
				checkErr(err)
			}
			for rows.Next() { //for loop through database table
				var oneMsg string
				err = rows.Scan(&oneMsg)
				checkErr(err)
				// fmt.Println("oneMsg1!", oneMsg)

				allMsg = append(allMsg, oneMsg)

			}
			if len(allMsg) != 0 {
				lastMsg = append(lastMsg, allMsg[len(allMsg)-1])

			} else {
				lastMsg = append(lastMsg, "0")
			}
		}
	}
	return lastMsg
}
func sendAllUsers(msg AllUsersOnline, wsConn *websocket.Conn) {
	err := wsConn.WriteJSON(msg)
	fmt.Println("All users online sended!")
	if err != nil {
		fmt.Printf("error sending message: %s\n", err.Error())
	}
}
func sendCheck(msg message, wsConn *websocket.Conn) {
	err := wsConn.WriteJSON(msg)
	fmt.Println("Session check sended!")
	if err != nil {
		fmt.Printf("error sending message: %s\n", err.Error())
	}
}
func saveMessage(data RecievedData, wsConn *websocket.Conn) {
	db, err := sql.Open("sqlite3", "./Database/forum.db")
	checkErr(err)
	stmt, err := db.Prepare(`INSERT INTO messages(content, created, user_1, user_2) values(?, datetime("now", "+180 minutes"),?,?)`)
	checkErr(err)
	stmt.Exec(data.Content, data.User1, data.User2)
	defer db.Close()

	var msg SendNotification
	msg.Sender = data.User1
	msg.Receiver = data.User2
	fmt.Println("Msg notification: ", msg)
	err = wsConn.WriteJSON(msg)
	if err != nil {
		fmt.Printf("error sending message: %s\n", err.Error())
	}
}
func sendMessage(data RecievedData, user1conn *websocket.Conn, user2conn *websocket.Conn) {

	var msg SendMsg
	timenow := time.Now()
	s := timenow.String()
	fmt.Println(s)
	created := timenow.Format("2006-01-02T15:04:05Z")
	time := created[11:16]
	msg = SendMsg{Content: data.Content, Sender: data.User1, Receiver: data.User2, Created: created, Date: "", Time: time}

	fmt.Println("msg sent back: ", msg)
	err := user1conn.WriteJSON(msg)
	if err != nil {
		fmt.Printf("error sending message1: %s\n", err.Error())
	}
	err2 := user2conn.WriteJSON(msg)
	if err2 != nil {
		fmt.Printf("error sending message2: %s\n", err.Error())
	}
}

var count = 1

func sendMessageHistory(user1 string, wsConn *websocket.Conn, activeChatUsers []string) {

	fmt.Println("User : ", user1)

	db, err := sql.Open("sqlite3", "./Database/forum.db")
	checkErr(err)

	var allMsg []SendMsg
	// var usr2 []SendMsg
	for _, user2 := range activeChatUsers {
		fmt.Println("User2(receiver) : ", user2)
		rows, err := db.Query("SELECT content, created, user_1, user_2 FROM messages WHERE user_1 = ? AND user_2 = ? OR user_1 = ? AND user_2 = ?", user1, user2, user2, user1)
		if err != nil {
			checkErr(err)
		}
		for rows.Next() { //for loop through database table
			oneMsg := SendMsg{}
			err = rows.Scan(&oneMsg.Content, &oneMsg.Created, &oneMsg.Sender, &oneMsg.Receiver)
			checkErr(err)
			oneMsg.Date = oneMsg.Created[:10]
			oneMsg.Time = oneMsg.Created[11:16]
			allMsg = append(allMsg, oneMsg)
		}
		var lastMsg []SendMsg
		for i := 1; i <= count; i++ {
			calc := len(allMsg) - i*10
			if calc > 0 {
				lastMsg = allMsg[calc:]

				err = wsConn.WriteJSON(lastMsg)
				if err != nil {
					fmt.Printf("error sending message: %s\n", err.Error())
				}

			}
		}
	}
	count++
}
