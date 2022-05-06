package database

import (
	"database/sql"
	"fmt"
	"log"

	"forum/likes"

	_ "github.com/mattn/go-sqlite3"
)

const dbName = "forum.db"

var PostIDInt int

func CreateDB() {
	var db *sql.DB
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}

	db.Exec("create table if not exists users (userID integer primary key, email text, username text, hash CHAR(60), usertype text, externalloginid text)")
	db.Exec(`create table if not exists post (
		postID integer primary key, 
		userID integer REFERENCES users(userID), 
		creationDate integer,
		postTitle CHAR(50),
		postContent CHAR(250), 
		postImages text, 
		edited integer);`)
	db.Exec("create table if not exists category (categoryID integer PRIMARY KEY, postID integer REFERENCES post(postID), categoryname text)")
	db.Exec(`create table if not exists comments (
		commentID integer primary key, 
		userID integer REFERENCES users(userID), 
		postID integer REFERENCES post(postID), 
		commentText CHAR(250), 
		edited integer, 
		creationDate integer);`)
	db.Exec(`create table if not exists likes (likeID integer PRIMARY KEY, 
		userID integer REFERENCES users(userID), postID integer REFERENCES post(postID), 
		commentID integer REFERENCES comments(commentID));`)
	db.Exec(`create table if not exists dislikes (dislikeID integer PRIMARY KEY, 
			userID integer REFERENCES users(userID), postID integer REFERENCES post(postID), 
			commentID integer REFERENCES comments(commentID));`)

	fmt.Println(likes.HomePostLikes(db))
}
