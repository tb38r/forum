package comments

import (
	"database/sql"
	"fmt"
)

type Comment struct {
	CommentID    int
	UserID       int
	PostID       int
	CreationDate int
	CommentText  string
	LikesID      int
	DislikesID   int
	Edited       bool
}

var db *sql.DB

var LastIns int64

func CreateComment(db *sql.DB, userID int, commentText string) {
	stmt, err := db.Prepare("INSERT INTO comments (userID, commentText, creationDate) VALUES (?,?, datetime('now', 'local time'))")
	if err != nil {
		fmt.Println("error preparing statement")
		return
	}

	result, _ := stmt.Exec(userID, commentText)

	rowsAff, _ := result.RowsAffected()
	LastIns, _ = result.LastInsertId()
	fmt.Println("rows affected: ", rowsAff)
	fmt.Println("last inserted: ", LastIns)
}
