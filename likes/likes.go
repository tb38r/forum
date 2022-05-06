package likes

import (
	"database/sql"
	"fmt"
)

type Like struct {
	LikeID    int
	UserID    int
	PostID    int
	CommentID int
}

var db *sql.DB

var LastIns int64

func LikeButton(db *sql.DB, userID int, postID int) {
	stmt, err := db.Prepare("INSERT INTO likes(userID, postID) VALUES(?, ?) ")
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

func GetLikeData(db *sql.DB, likeID int) Like {
	row := db.QueryRow("SELECT * FROM like WHERE likeID = ?;", likeID)
	var like Like
	err := row.Scan(&like.PostID, &like.UserID, &like.PostID, &like.CommentID)
	if err != nil {
		fmt.Println(err)
	}
	return like
}
