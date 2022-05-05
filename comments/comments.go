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
	rows, err := db.Query("SELECT commentText FROM comments INNER JOIN post ON post.postID = comments.postID WHERE post.postID = ?;", postID)
	if err != nil {
		fmt.Println(err)
	}
	comment := []Comment{}
	defer rows.Close()
	for rows.Next() {
		var c Comment
		err2 := rows.Scan(&c.CommentText)
		comment = append(comment, c)
		if err2 != nil {
			fmt.Println(err2)
		}
	}
	return comment
}
