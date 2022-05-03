package web

import (
	"database/sql"
	"fmt"
	"forum/comments"
	"net/http"
	"strconv"
)

var CUserIdint int
var ContentComment string

//var CPostIdint int

func (s *myServer) CreateCommentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		userID := r.URL.Query().Get("userid")

		CUserIdint, _ = strconv.Atoi(userID)

		Tpl.ExecuteTemplate(w, "createcomment.html", nil)

	}
}

func (s *myServer) StoreCommentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()
		ContentComment = r.FormValue("content")
		// postID := r.URL.Query().Get("postid")
		// CPostIdint, _ = strconv.Atoi(postID)

		comments.CreateComment(s.Db, GuserId, PostIDInt, ContentComment)

		var commentData comments.Comment

		commentData.CommentText = ContentComment
		commentData.PostID = PostIDInt
		commentData.UserID = GuserId

		fmt.Println("comment data check: ---> ", commentData.CommentText)
		fmt.Println("comment post id check: ---> ", commentData.PostID)
		fmt.Println("comment user id check: ---> ", commentData.UserID)

		fmt.Println("content: ", ContentComment)
		Tpl.ExecuteTemplate(w, "storecomment.html", commentData.CommentText)

	}
}

func (s *myServer) ShowCommentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		// had to open the database here as it wasnt picking the correct post everytime without this.
		s.Db, _ = sql.Open("sqlite3", "forum.db")
		// get the postId and display the post and its contents
		postID := r.URL.Query().Get("postid")
		PostIDInt, _ = strconv.Atoi(postID)

		var commentData comments.Comment

		if PostIDInt == commentData.PostID {
			commentData.CommentText = ContentComment
			commentData.PostID = PostIDInt
			commentData.UserID = GuserId
		}
		fmt.Println("comment data check: ---> ", commentData.CommentText)
		fmt.Println("comment post id check: ---> ", commentData.PostID)

		Tpl.ExecuteTemplate(w, "showcomment.html", commentData)
	}
}
