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

		SPostID := strconv.Itoa(PostIDInt)

		if !UserDisliked(s.Db) {
			dislikes.DislikeButton(s.Db, GuserId, PostIDInt)
			fmt.Println("Dislike added to database----------------------")
			http.Redirect(w, r, "/showpost/?postid="+SPostID, http.StatusSeeOther)
		} else {
			fmt.Println("Disike not added to database-----------------------")
			http.Redirect(w, r, "/showpost/?postid="+SPostID, http.StatusSeeOther)
		}
	}
}

func UserDisliked(db *sql.DB) bool {
	// check if post already liked by user
	userStmt := "SELECT userID FROM dislikes WHERE postID = ?"
	row := db.QueryRow(userStmt, PostIDInt)

	var uID string
	var pID string
	err := row.Scan(&uID, &pID)
	if err != sql.ErrNoRows {
		return true
	}
	fmt.Println("User has already disliked this post", err)
	return false
}
