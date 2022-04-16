package main

import "time"

type user struct {
	UserName string
	Password []byte
	Email    string
}

type createdposts struct {
	PostID string
	Title  string
}
type profilepage struct {
	Createdposts []createdposts
	User         *user
	Likedposts   []likedposts
}
type likedposts struct {
	PostID string
	Title  string
}

type allpost struct {
	Post     []post
	Category []category
}

type mainstr struct {
	User    *user
	Allpost *allpost
}

type category struct {
	Category []string
	IDc      int
}

type postdata struct {
	Post           *post
	User           *user
	Category       *category
	Comarr         *comarr
	Postlike       *postlike
	Commentlikearr *commentlikearr
}
type postlike struct {
	Like    int
	Dislike int
}
type commentlike struct {
	Like    int
	Dislike int
}
type comarr struct {
	Comment []comment
}
type commentlikearr struct {
	Comlike []commentlike
}
type session struct {
	un           string
	lastActivity time.Time
}
