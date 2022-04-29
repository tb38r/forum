package posts

import (
	"database/sql"
	"fmt"
	"log"
)

type Post struct {
	PostID       int
	UserID       int
	CommentID    int
	CategoryID   int
	CreationDate int
	PostTitle    string
	PostText     string
	PostImage    string
	LikesID      int
	DislikesID   int
	Edited       bool
}

var db *sql.DB

// type s *web.Server

var LastIns int64

func CreatePosts(db *sql.DB, userID int, title string, content string) {
	stmt, err := db.Prepare("INSERT INTO post (userID, postTitle, postContent, creationDate) VALUES (?, ?, ?, datetime('now', 'localtime'))")
	if err != nil {
		fmt.Println("error preparing statement:", err)
		return
	}
	// defer stmt.Close()
	result, _ := stmt.Exec(userID, title, content)

	// checking if the result has been added and the last inserted row
	rowsAff, _ := result.RowsAffected()
	LastIns, _ = result.LastInsertId()
	fmt.Println("rows affected:", rowsAff)
	fmt.Println("last inserted:", LastIns)
}

// function that gets all the post titles and returns a slice of string
func GetAllPostTitles(db *sql.DB) []string {
	rows, err := db.Query("SELECT postTitle FROM post")

	if err != nil {
		fmt.Println(err)
	}
	AllpostTitles := []string{}
	var postTitle string
	defer rows.Close()
	for rows.Next() {
		err2 := rows.Scan(&postTitle)
		if err2 != nil {
			log.Fatal(err2)
		}
		AllpostTitles = append(AllpostTitles, postTitle)
	}
	return AllpostTitles
}
