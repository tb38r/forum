package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const dbName = "forum.db"

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
		commentID integer REFERENCES comment(commentID), 
		categoryID integer REFERENCES category(categoryID), 
		creationDate integer,
		postTitle CHAR(50),
		postContent CHAR(250), 
		postImages text, 
		likeID integer REFERENCES like(likeID), 
		dislikeID integer REFERENCES dislike(dislikeID), 
		edited integer);`)
	// db.Exec("create table if not exists category (categoryID integer PRIMARY KEY, postID integer REFERENCES post(postID), categoryname text)")
	db.Exec(`create table if not exists comments (commentID integer PRIMARY KEY, userID integer REFERENCES users(userID), postID integer post(postID);`)
	db.Exec(`create table if not exists likes (likeID integer PRIMARY KEY, userID integer REFERENCES users(userID), postID integer post(postID), commentID integer REFERENCES comments(commentID));`)
	// db.Exec("create table if not exists dislikes (DislikeID integer primary key, UserID integer foreign key, CommentID integer foreign key, PostID integer foreign key)")
}
