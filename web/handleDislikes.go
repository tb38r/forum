package web

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"forum/comments"
	"forum/dislikes"
	"forum/likes"
)

// var CUserIdint int

func (s *myServer) DislikeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("dislike handler running")

		SPostID := strconv.Itoa(PostIDInt)

		if !UserDisliked(s.Db) || UserLiked(s.Db) {
			dislikes.DislikeButton(s.Db, GuserId, PostIDInt)
			likes.DeleteLike(s.Db, GuserId, PostIDInt)
			fmt.Println("Dislike added to database----------------------")
			http.Redirect(w, r, "/showpost/?postid="+SPostID, http.StatusSeeOther)
		} else {
			fmt.Println("Disike not added to database-----------------------")
			http.Redirect(w, r, "/showpost/?postid="+SPostID, http.StatusSeeOther)
		}
	}
}

func (s *myServer) CommentDislikeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("comment dislike handler running")

		SPostID := strconv.Itoa(PostIDInt)

		CommentId := comments.GetCommentID(s.Db)

		if !CommentUserDisliked(s.Db) || UserLiked(s.Db) {
			dislikes.CommentDislikeButton(s.Db, GuserId, CommentId)
			likes.DeleteCommentLike(s.Db, GuserId, CommentId)
			fmt.Println(" Comment Dislike added to database----------------------")
			http.Redirect(w, r, "/showpost/?postid="+SPostID, http.StatusSeeOther)
		} else {
			fmt.Println("Comment Disike not added to database-----------------------")
			http.Redirect(w, r, "/showpost/?postid="+SPostID, http.StatusSeeOther)
		}
	}
}

func UserDisliked(db *sql.DB) bool {
	// check if post already disliked by user
	userStmt := "SELECT userID FROM dislikes WHERE userID = ? AND postID = ?"
	row := db.QueryRow(userStmt, GuserId, PostIDInt)

	var uID string
	var pID string
	err := row.Scan(&uID, &pID)
	if err != sql.ErrNoRows {
		return true
	}
	fmt.Println("User has already disliked this post", err)
	return false
}

func CommentUserDisliked(db *sql.DB) bool {
	// check if comment already disliked by user
	userStmt := "SELECT userID FROM dislikes WHERE userID = ? AND commentID = ?"
	row := db.QueryRow(userStmt, GuserId, CommentId)

	var uID string
	var cID string
	err := row.Scan(&uID, &cID)
	if err != sql.ErrNoRows {
		return true
	}
	fmt.Println("User has already disliked this comment", err)
	return false
}
