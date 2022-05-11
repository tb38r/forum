package web

import (
	"database/sql"
	"forum/comments"
	"forum/posts"
	"forum/users"
	"net/http"
)

type ActivityPage struct {
	Posts    []posts.Post
	Likes    []posts.Post
	Dislikes []posts.Post
	Comments []posts.Post
}

func (s *myServer) ShowActivityPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// had to open the database here as it wasnt picking the correct post everytime without this.
		s.Db, _ = sql.Open("sqlite3", "forum.db")
		// get the postId and display the post and its contents

		// postID := r.URL.Query().Get("postid")
		// PostIDInt, _ = strconv.Atoi(postID)

		data := ActivityPage{Posts: posts.GetPostData(s.Db, PostIDInt), Comments: comments.GetCommentData(s.Db, PostIDInt), Loggedin: users.AlreadyLoggedIn(r)}

		Tpl.ExecuteTemplate(w, "showpost.html", data)

	}
}
