package web

import (
	"fmt"
	"forum/posts"
	"net/http"
)

type ActivityPage struct {
	Posts    []posts.Post
	Likes    []posts.Post
	Dislikes []posts.Post
	Comments []posts.Post
}

func (s *myServer) ActivityPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "PLACEHOLDER")

	}
}

// had to open the database here as it wasnt picking the correct post everytime without this.
// s.Db, _ = sql.Open("sqlite3", "forum.db")
// // get the postId and display the post and its contents

// // postID := r.URL.Query().Get("postid")
// // PostIDInt, _ = strconv.Atoi(postID)

// data := ActivityPage{Posts: posts.GetPostData(s.Db, PostIDInt)}

// Tpl.ExecuteTemplate(w, "showpost.html", data)
