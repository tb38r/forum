package web

import (
	"fmt"
	"net/http"
	//"strconv"
)

var CUserIdint int

func (s *Server) LikeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("like handler running")
		
		r.ParseForm()
		
		// postID := r.URL.Query().Get("postid")
		// PostIdint, _ = strconv.Atoi(postID)

		like := r.FormValue("like")
		fmt.Println("what is this", like)
		tpl.ExecuteTemplate(w, "likes.html", nil)
	}
}
