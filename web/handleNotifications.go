package web

import (
	"database/sql"
	"log"
)

//	 JOIN post ON post.postID = comments.postID 

// rows, err := db.Query(`SELECT users.username, post.postTitle
// FROM comments, users, post
//  JOIN comments ON comments.userID = users.userID
//  WHERE comments.creatorID = ?
//  AND post.postID = comments.creatorID 
//  AND comments.notified = ?
// ;`, GuserId, 0)
// if err != nil {
// 	log.Fatal("Web CommentUsername Error:", err)

// }

//return userID & postID where comment was made where notifed = 0
func CommentUsername(db *sql.DB) []map[string]int {
	rows, err := db.Query(`SELECT users.username, comments.postID
	FROM users, comments
	 WHERE comments.creatorID = ?
	 AND comments.notified = ?
	 AND comments.userID = users.userID
	;`, GuserId, 0)
	if err != nil {
		log.Fatal("Web CommentUsername Error:", err)

	}

	var CommentNotification []map[string]int

	var username string
	var postid int

	defer rows.Close()
	for rows.Next() {
		minimap := make (map[string]int)
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
