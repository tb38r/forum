package web

import (
	"fmt"
	"net/http"

	"forum/likes"
	//"strconv"
)

// var CUserIdint int

func (s *myServer) LikeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("like handler running")

		r.ParseForm()

		like := r.FormValue("like")

	
	    likes.LikeButton(s.Db, GuserId, PostIDInt)
		

		fmt.Println("what is this", like)
		Tpl.ExecuteTemplate(w, "likes.html", nil)
	}
}
