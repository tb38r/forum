package comments

import (
	"database/sql"
	"fmt"
	"forum/dislikes"
	"forum/likes"
	"log"
)

type Comment struct {
	CommentID       int
	UserID          int
	PostID          int
	CreationDate    string
	CommentText     string
	CommentUserName string
	Likes           int
	Dislikes        int
	Edited          bool
}

var db *sql.DB

var LastIns int64

func CreateComment(db *sql.DB, userID int, postID int, commentText string) {
	stmt, err := db.Prepare("INSERT INTO comments (userID, postID, commentText, creationDate) VALUES (?, ?, ?, strftime('%H:%M %d/%m/%Y','now', 'localtime'))")

	if err != nil {
		fmt.Println("error preparing statement")
		return
	}

	result, _ := stmt.Exec(userID, postID, commentText)

	rowsAff, _ := result.RowsAffected()
	LastIns, _ = result.LastInsertId()
	fmt.Println("rows affected: ", rowsAff)
	fmt.Println("last inserted: ", LastIns)
}

func GetCommentText(db *sql.DB) map[int]string {
	rows, err := db.Query("SELECT postID, commentText FROM comments")

	if err != nil {
		fmt.Println("gct error 1", err)
	}
	// AllpostTitles := []string{}
	CommentText := make(map[int]string)

	var postID int
	var commentText string
	defer rows.Close()
	for rows.Next() {
		err2 := rows.Scan(&postID, &commentText)
		if err2 != nil {
			log.Fatal("gct err 2", err2)
		}
		// AllpostTitles = append(AllpostTitles, postID, postTitle)
		CommentText[postID] = commentText
	}
	return CommentText
}

func GetCommentData(db *sql.DB, postID int) []Comment {
	rows, err := db.Query(`SELECT commentID, commentText, comments.creationDate as cmntDate, users.username 
	FROM comments 
	INNER JOIN post ON post.postID = comments.postID 
	INNER JOIN users ON users.userID = comments.userID 
	WHERE post.postID = ?;`, postID)
	if err != nil {
		fmt.Println(err)
	}
	comment := []Comment{}
	defer rows.Close()
	for rows.Next() {
		var c Comment
		err2 := rows.Scan(&c.CommentID, &c.CommentText, &c.CreationDate, &c.CommentUserName)
		c.Likes = likes.GetCommentLikes(db, c.CommentID)
		c.Dislikes = dislikes.GetCommentDislikes(db, c.CommentID)
		comment = append(comment, c)
		if err2 != nil {
			fmt.Println(err2)
		}
	}
	return comment
}

func GetCommentID(db *sql.DB, postID int) int {
	rows, err := db.Query(`SELECT comments.commentID from comments where comments.postID = ?;`, postID)
	if err != nil {
		fmt.Println(err)
	}
	var cmt Comment
	defer rows.Close()
	for rows.Next() {
		var c Comment
		err2 := rows.Scan(&c.CommentID)
		cmt = c
		if err2 != nil {
			fmt.Println(err2)
		}
	}
	return cmt.CommentID
}

func (c Comment) GetCID() int {
	sqlStatement := `SELECT comments.commentID FROM comments;`

	var id int

	// Replace 3 with an ID from your database or another random
	// value to test the no rows use case.
	row := db.QueryRow(sqlStatement)
	switch err := row.Scan(&id); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(id)
	default:
		panic(err)
	}
	c.CommentID = id
	return c.CommentID
}
