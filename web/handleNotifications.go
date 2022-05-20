package web

import (
	"database/sql"
	"log"
)


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

//return userID & postID where comment was made where notifed = 0
func CommentUsername(db *sql.DB) []map[string]string {
	rows, err := db.Query(`SELECT users.username, post.postTitle
	FROM users, comments, post 
	 WHERE comments.creatorID = ?
	 AND comments.notified = ?
	 AND comments.userID = users.userID
	 AND comments.postID = post.postID
	;`, GuserId, 0)
	if err != nil {
		log.Fatal("Web CommentUsername Error:", err)

	}

	var CommentNotification []map[string]string

	var username string
	var postid string

	defer rows.Close()
	for rows.Next() {
		minimap := make (map[string]string)
		err2 := rows.Scan(&username, &postid)
		if err2 != nil {
			log.Fatal("Web CommentUsername Error:", err2)
		}
		// AllpostTitles = append(AllpostTitles, postID, postTitle)
		minimap[username] = postid
		CommentNotification = append(CommentNotification, minimap)
	}
	return CommentNotification

}
