package web

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"forum/dislikes"
	"forum/likes"
)

// var CUserIdint int

func (s *myServer) LikeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("like handler running")

		SPostID := strconv.Itoa(PostIDInt)

		if !UserLiked(s.Db) || UserDisliked(s.Db) {
			likes.LikeButton(s.Db, GuserId, PostIDInt)
			dislikes.DeleteDislike(s.Db, GuserId, PostIDInt)
			fmt.Println("Like added to database----------------------")
			http.Redirect(w, r, "/showpost/?postid="+SPostID, http.StatusSeeOther)
		} else {
			fmt.Println("Like not added to database-----------------------")
			http.Redirect(w, r, "/showpost/?postid="+SPostID, http.StatusSeeOther)
		}
	}
}

func (s *myServer) CommentLikeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("comment like handler running")

		SPostID := strconv.Itoa(PostIDInt)

		if !CommentUserLiked(s.Db) || CommentUserDisliked(s.Db) {
			likes.CommentLikeButton(s.Db, GuserId, CommentId)
			dislikes.DeleteCommentDislike(s.Db, GuserId, CommentId)
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
	userStmt := "SELECT userID FROM likes WHERE userID = ? AND postID = ?"
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

func CommentUserLiked(db *sql.DB) bool {
	// check if post already liked by user
	userStmt := "SELECT userID FROM likes WHERE userID = ? AND commentID = ?"
	row := db.QueryRow(userStmt, GuserId, CommentId)
	fmt.Println("----------***********************----------- checking if commentData.CommentID works", CommentData.CommentID)

	var uID string
	var cID string
	err := row.Scan(&uID, &cID)
	if err != sql.ErrNoRows {
		return true
	}
	fmt.Println("User has already liked this comment", err)
	return false
}
