package posts

import (
	"database/sql"
	"fmt"
	"log"
)

type Post struct {
	PostID       int
	UserID       int
	CreationDate string
	PostTitle    string
	PostContent  string
	PostImage    string
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

// function that gets all the post titles and returns a slice of string, also now getting the postId
func GetAllPostTitles(db *sql.DB) map[int]string {
	rows, err := db.Query("SELECT postID, postTitle FROM post")

	if err != nil {
		fmt.Println(err)
	}
	// AllpostTitles := []string{}
	AllpostTitles := make(map[int]string)

	var postID int
	var postTitle string
	defer rows.Close()
	for rows.Next() {
		err2 := rows.Scan(&postID, &postTitle)
		if err2 != nil {
			log.Fatal(err2)
		}
		// AllpostTitles = append(AllpostTitles, postID, postTitle)
		AllpostTitles[postID] = postTitle
	}
	return AllpostTitles
}

// getting the data from one post, and storing the values in the post struct
func GetPostData(db *sql.DB, postID int) Post {
	row := db.QueryRow("SELECT * FROM post WHERE postID = ?;", postID)
	var post Post
	err := row.Scan(&post.PostID, &post.UserID, &post.CreationDate, &post.PostTitle, &post.PostContent, &post.PostImage, &post.Edited)
	if err != nil {
		fmt.Println(err)
	}
	return post
}
