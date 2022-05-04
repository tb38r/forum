package dislikes

import (
	"database/sql"
	"fmt"
)

type Disike struct {
	DislikeID    int
	UserID    int
	PostID    int
	CommentID int
}

var db *sql.DB

var LastIns int64

func DislikeButton(db *sql.DB, userID int, postID int) {
	stmt, err := db.Prepare("INSERT INTO dislikes(userID, postID) VALUES(?, ?) ")
	if err != nil {
		fmt.Println("error preparing statement:", err)
		return
	}

	result, _ := stmt.Exec(userID, postID)

	rowsAff, _ := result.RowsAffected()
	LastIns, _ = result.LastInsertId()
	fmt.Println("rows affected:", rowsAff)
	fmt.Println("last inserted", LastIns)
}
