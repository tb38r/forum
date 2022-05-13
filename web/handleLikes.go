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
var CMNTID int

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
		r.ParseForm()
		for k, v := range r.Form {
			fmt.Println(k, v)
		}

		CMNTID, _ = strconv.Atoi(r.FormValue("commentlike"))
		fmt.Println("comment like handler running")

		SPostID := strconv.Itoa(PostIDInt)
		//cid := cmt.GetCID()
		//fmt.Println("checking if id is in struct", cmt.CommentID)
		fmt.Println("Checking what CommentID is in like handler", CommentId)
		//fmt.Println("Checking what cID method is in like handler", cid)

		if !CommentUserLiked(s.Db) || CommentUserDisliked(s.Db) {
			likes.CommentLikeButton(s.Db, GuserId, CMNTID)
			dislikes.DeleteCommentDislike(s.Db, GuserId, CMNTID)
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
		fmt.Println("User has already liked this post", err)
		return true
	}
	return false
}

func CommentUserLiked(db *sql.DB) bool {
	// check if post already liked by user
	userStmt := "SELECT userID FROM likes WHERE userID = ? AND commentID = ?"

	row := db.QueryRow(userStmt, GuserId, CMNTID)
	fmt.Println("----------***********************----------- checking if commentData.CommentID works", CMNTID)

	var uID string
	var cID string
	err := row.Scan(&uID, &cID)
	if err != sql.ErrNoRows {
		fmt.Println("User has already liked this comment", err)
		return true
	}
	return false
}
