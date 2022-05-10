package web

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"forum/dislikes"
)

// var CUserIdint int

func (s *myServer) DislikeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("dislike handler running")

		if !userDisliked(s.Db) {
			dislikes.DislikeButton(s.Db, GuserId, PostIDInt)
		}
		fmt.Println(PostIDInt)

		SPostID := strconv.Itoa(PostIDInt)

		http.Redirect(w, r, "/showpost/?postid="+SPostID, http.StatusSeeOther)
	}
}

func userDisliked(db *sql.DB) bool {
	// check if post already liked by user
	userStmt := "SELECT userID FROM dislikes WHERE postID = ?"
	row := db.QueryRow(userStmt, PostIDInt)
	var uID string
	err := row.Scan(&uID)
	if err != sql.ErrNoRows {
		fmt.Println("User has already disliked this post, err:", err)
		return true
	}
	return false
}
