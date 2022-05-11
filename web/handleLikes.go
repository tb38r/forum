package web

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"forum/likes"
)

// var CUserIdint int

func (s *myServer) LikeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("like handler running")

		SPostID := strconv.Itoa(PostIDInt)

		if !UserLiked(s.Db) {
			likes.LikeButton(s.Db, GuserId, PostIDInt)
			fmt.Println("Like added to database----------------------")
			http.Redirect(w, r, "/showpost/?postid="+SPostID, http.StatusSeeOther)
		} else {
			fmt.Println("Like not added to database-----------------------")
			http.Redirect(w, r, "/showpost/?postid="+SPostID, http.StatusSeeOther)
		}
	}
}

func UserLiked(db *sql.DB) bool {
	// check if post already liked by user
	userStmt := "SELECT userID FROM likes WHERE userID = 1 AND postID = 4"
	row := db.QueryRow(userStmt, GuserId, PostIDInt)

	var uID string 
	var pID string
	err := row.Scan(&uID, &pID)
	if err != sql.ErrNoRows {
		return true
	}
	fmt.Println("User has already liked this post", err)
	return false
}
