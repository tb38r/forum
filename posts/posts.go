package posts

import "database/sql"

type Post struct {
	PostID       int
	UserID       int
	CommentID    int
	CategoryID   int
	CreationDate int
	PostText     string
	PostImage    string
	LikesID      int
	DislikesID   int
	Edited       bool
}

var db *sql.DB

// func createPosts(db *sql.DB)
