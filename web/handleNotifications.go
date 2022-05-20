package web

import (
	"database/sql"
	"log"
)

type CommNotify struct {
	Username  string
	PostID    int
	PostTitle string
}

// //return userID & postID where comment was made where notifed = 0
// func CommentUsername(db *sql.DB) []map[string]int {
// 	rows, err := db.Query(`SELECT users.username, comments.postID
// 	FROM users, comments
// 	 WHERE comments.creatorID = ?
// 	 AND comments.notified = ?
// 	 AND comments.userID = users.userID
// 	;`, GuserId, 0)
// 	if err != nil {
// 		log.Fatal("Web CommentUsername Error:", err)

// 	}

// 	var CommentNotification []map[string]int

// 	var username string
// 	var postid int

// 	defer rows.Close()
// 	for rows.Next() {
// 		minimap := make (map[string]int)
// 		err2 := rows.Scan(&username, &postid)
// 		if err2 != nil {
// 			log.Fatal("Web CommentUsername Error:", err2)
// 		}
// 		// AllpostTitles = append(AllpostTitles, postID, postTitle)
// 		minimap[username] = postid
// 		CommentNotification = append(CommentNotification, minimap)
// 	}
// 	return CommentNotification

// }

//return user & postid where comment was made where notified = 0
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

func ResetCommentNotified(db *sql.DB) {
	db.Exec(`UPDATE comments
	SET comments.notified = 1
	WHERE comments.creatorID = ?
	;`, GuserId)

}
