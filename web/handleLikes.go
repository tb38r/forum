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

		//if !userLiked(s.Db, GuserId, PostIDInt) {
			likes.LikeButton(s.Db, GuserId, PostIDInt)
		//}
		fmt.Println(PostIDInt)

		http.Redirect(w, r, "/showpost/?postid="+SPostID, http.StatusSeeOther)
	}
}

func userLiked(db *sql.DB, userID int, postID int) bool {
	// check if post already liked by user
	err := db.QueryRow("SELECT userID FROM likes WHERE userID = ? AND postID = ?", GuserId, PostIDInt)

	fmt.Println("user liked------------------------------", err)

	return err != nil
}
