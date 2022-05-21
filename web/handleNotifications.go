package web

import (
	"database/sql"
	"fmt"
	"log"
)

type CommNotify struct {
	Username  string
	PostID    int
	PostTitle string
}

//return user, postid & posttitle where comment was made where notified = 0
func CommentNotify(db *sql.DB) []CommNotify {
	rows, err := db.Query(`SELECT users.username, comments.postID, post.postTitle
	FROM users, comments, post
	 WHERE comments.creatorID = ?
	 AND comments.userID != ?
	 AND comments.notified = ?
	 AND comments.userID = users.userID
	 AND comments.postID = post.postID

	;`, GuserId, GuserId, 0)
	if err != nil {
		log.Fatal("Web CommentUsername Error:", err)

	}

	CommentNotification := []CommNotify{}

	var username string
	var postid int
	var pTitle string

	defer rows.Close()
	for rows.Next() {
		ministruct := CommNotify{}
		err2 := rows.Scan(&username, &postid, &pTitle)
		if err2 != nil {
			log.Fatal("Web CommentUsername Error:", err2)
		}
		ministruct.Username = username
		ministruct.PostID = postid
		ministruct.PostTitle = pTitle

		CommentNotification = append(CommentNotification, ministruct)
	}

	return CommentNotification

}

//return user, postid & posttitle where comment was made where notified = 0
func LikesNotify(db *sql.DB) []CommNotify {
	rows, err := db.Query(`SELECT users.username, likes.postID, post.postTitle
	FROM users, likes, post
	 WHERE likes.creatorID = ?
	 AND likes.userID != ?
	 AND likes.notified = ?
	 AND likes.userID = users.userID
	 AND likes.postID = post.postID

	;`, GuserId, GuserId, 0)
	if err != nil {
		log.Fatal("Web LikeNotify Error:", err)

	}

	LikeNotification := []CommNotify{}

	var username string
	var postid int
	var pTitle string

	defer rows.Close()
	for rows.Next() {
		ministruct := CommNotify{}
		err2 := rows.Scan(&username, &postid, &pTitle)
		if err2 != nil {
			log.Fatal("LikesNotify Error:", err2)
		}
		ministruct.Username = username
		ministruct.PostID = postid
		ministruct.PostTitle = pTitle

		LikeNotification = append(LikeNotification, ministruct)
	}

	return LikeNotification

}

func ResetCommentNotified(db *sql.DB) {
	stmt, err := db.Prepare(`UPDATE comments
	SET notified = ?
	WHERE creatorID = ?
	;`)
	defer stmt.Close()
	if err != nil {
		log.Fatal("ResetComment 1:", err)
	}

	res, err2 := stmt.Exec(1, GuserId)
	if err2 != nil {
		log.Fatal("ResetComment 2:", err)
	}

	affected, err3 := res.RowsAffected()
	if err3 != nil {
		log.Fatal("ResetComment 3:", err)
	}

	fmt.Println(affected)

}

func ResetLikesNotified(db *sql.DB) {
	stmt, err := db.Prepare(`UPDATE likes
	SET notified = ?
	WHERE creatorID = ?
	;`)
	defer stmt.Close()
	if err != nil {
		log.Fatal("ResetLikes 1:", err)
	}

	res, err2 := stmt.Exec(1, GuserId)
	if err2 != nil {
		log.Fatal("ResetLikes 2:", err)
	}

	affected, err3 := res.RowsAffected()
	if err3 != nil {
		log.Fatal("ResetLikes 3:", err)
	}

	fmt.Println("Likes Affected:", affected)

}
