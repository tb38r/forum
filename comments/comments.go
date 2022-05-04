package comments

import (
	"database/sql"
	"fmt"
	"log"
)

type Comment struct {
	CommentID    int
	UserID       int
	PostID       int
	CreationDate int
	CommentText  string
	//LikesID      int
	//DislikesID   int
	//Edited       bool
}

var db *sql.DB

var LastIns int64

func CreateComment(db *sql.DB, userID int, postID int, commentText string) {
	stmt, err := db.Prepare("INSERT INTO comments (userID, postID, commentText, creationDate) VALUES (?,?,?, datetime('now', 'local time'))")
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
		fmt.Println(err)
	}
	// AllpostTitles := []string{}
	CommentText := make(map[int]string)

	var postID int
	var commentText string
	defer rows.Close()
	for rows.Next() {
		err2 := rows.Scan(&postID, &commentText)
		if err2 != nil {
			log.Fatal(err2)
		}
		// AllpostTitles = append(AllpostTitles, postID, postTitle)
		CommentText[postID] = commentText
	}
	return CommentText
}

func GetCommentData(db *sql.DB, postID int) Comment {
	row := db.QueryRow("SELECT * FROM comments WHERE postID = ?;", postID)
	var comment Comment
	err := row.Scan(&comment.PostID, &comment.UserID, &comment.CreationDate, &comment.CommentText)
	if err != nil {
		fmt.Println(err)
	}
	return comment
}
