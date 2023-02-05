# Real Time Forum

Includes: 
* Registration and Login
* User authentification, sessions and cookies 
* Adding posts and comments and saving into database using SQLite
* Private messages using Go Websockets

* Front-end build using Vue.js framework
* Back-end build using Go

Project requirements are found here `https://01.kood.tech/git/root/public/src/branch/master/subjects/real-time-forum`


## How to run

- Clone the repo
- Go to ./real-time-forum/app and install packages for Vuejs using `npm install` and `npm run build`.
- Run Vuejs server ./real-time-forum/app using `npm run dev`
- Run Golang server ./real-time-forum/server using `go run .`
- Vuejs server will run at `http://localhost:3000/` and Golang's at `http://localhost:8090/`
- Open Vue server or `http://localhost:3000/` to see Real-time-forum!
