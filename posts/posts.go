package posts

import (
	"database/sql"
	"fmt"
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

type HomepagePosts struct {
	PostID       int
	PostTitle    string
	PostUsername string
	CreationDate string
}

var db *sql.DB

// type s *web.Server

var LastIns int64

func CreatePosts(db *sql.DB, userID int, title string, content string) {
	stmt, err := db.Prepare("INSERT INTO post (userID, postTitle, postContent, creationDate) VALUES (?, ?, ?, strftime('%H:%M %d/%m/%Y','now','localtime'))")

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

// Get all the data needed for the hompage
func GetHomepageData(db *sql.DB) []HomepagePosts {
	rows, err := db.Query("SELECT postID, postTitle, username, creationDate FROM post INNER JOIN users ON users.userID = post.userID;")
	if err != nil {
		fmt.Println(err)
	}
	postdata := []HomepagePosts{}
	defer rows.Close()
	for rows.Next() {
		var p HomepagePosts
		err2 := rows.Scan(&p.PostID, &p.PostTitle, &p.PostUsername, &p.CreationDate)
		postdata = append(postdata, p)
		if err2 != nil {
			fmt.Println(err2)
		}
	}
	return postdata

}

// Gets data based on user's filter choice (currently displays user's created posts, //TODO : Return Liked Posts)
func FilterHomepageData(db *sql.DB, userID int) []HomepagePosts {
	rows, err := db.Query("SELECT postID, postTitle, username, creationDate FROM post INNER JOIN users ON users.userID =  post.userID WHERE users.userID = ?;", userID)
	if err != nil {
		fmt.Println(err)
	}
	postdata := []HomepagePosts{}
	defer rows.Close()
	for rows.Next() {
		var p HomepagePosts
		err2 := rows.Scan(&p.PostID, &p.PostTitle, &p.PostUsername, &p.CreationDate)
		postdata = append(postdata, p)
		if err2 != nil {
			fmt.Println(err2)
		}
	}
	return postdata

}

// getting the data from one post, and storing the values in the post struct
func GetPostData(db *sql.DB, postID int) []Post {
	rows, err := db.Query("SELECT * FROM post WHERE postID = ?;", postID)
	if err != nil {
		fmt.Println(err)
	}
	postdata := []Post{}
	defer rows.Close()
	for rows.Next() {
		var p Post
		err2 := rows.Scan(&p.PostID, &p.UserID, &p.CreationDate, &p.PostTitle, &p.PostContent, &p.PostImage, &p.Edited)
		postdata = append(postdata, p)
		if err2 != nil {
			fmt.Println(err2)
		}
	}
	return postdata
}
