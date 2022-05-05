package web

import (
	"fmt"
	"forum/comments"
	"net/http"
	"strconv"
)

var CommentData comments.Comment
var CUserIdint int
var ContentComment string
var CommentMap = make(map[string]int)

//var CPostIdint int

func (s *myServer) CreateCommentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		userID := r.URL.Query().Get("userid")

		PostIDInt, _ = strconv.Atoi(userID)

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
		CommentMap[ContentComment] = PostIDInt

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

// func (s *myServer) ShowCommentHandler() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		r.ParseForm()
// 		// had to open the database here as it wasnt picking the correct post everytime without this.
// 		s.Db, _ = sql.Open("sqlite3", "forum.db")
// 		// get the postId and display the post and its contents
// 		postID := r.URL.Query().Get("postid")
// 		PostIDInt, _ = strconv.Atoi(postID)

// 		if PostIDInt == CommentData.PostID {
// 			CommentData.CommentText = ContentComment
// 			CommentData.PostID = PostIDInt
// 			CommentData.UserID = GuserId
// 		}
// 		fmt.Println("comment data check: ---> ", CommentData.CommentText)
// 		fmt.Println("comment post id check: ---> ", CommentData.PostID)

// 		for i := CommentData.CommentID; i >= 0; i-- {
// 			fmt.Fprintln(w, "<h1>"+CommentData.CommentText+"</h1>")
// 			//fmt.Fprintln(w, "The users id is ---> ", CommentData.UserID)
// 			fmt.Fprintln(w, "<h3>"+"<pre>"+users.CurrentUser+"</pre>"+"<h3>")
// 			//fmt.Fprintln(w, "The post id is ---> ", CommentData.PostID)
// 		}
// 		//Tpl.ExecuteTemplate(w, "showcomment.html", comments.GetCommentData(s.Db, PostIDInt))
// 	}
// }
