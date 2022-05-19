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

//return userID & postID where comment was made
func CommentUsername(db *sql.DB) map[int]int {
	rows, err := db.Query(`SELECT comments.userID, comments.postID
	FROM comments,  
	 JOIN comments ON comments.userID = users.userID
	 WHERE comments.creatorID = ?
	 JOIN posts ON post.postID = comments.postID 
	 WHERE comments.creatorID = ?
	 AND comments.notified = ?
	;`, GuserId, GuserId, 0)
	if err != nil {
		log.Fatal("Web CommentUsername Error:", err)

	}

	CommentNotification := make(map[string]string)

	var username string
	var title string

	defer rows.Close()
	for rows.Next() {
		err2 := rows.Scan(&username, &title)
		if err2 != nil {
			log.Fatal("Web CommentUsername Error:", err2)
		}
		// AllpostTitles = append(AllpostTitles, postID, postTitle)
		CommentNotification[username] = title
	}
	return CommentNotification

}
