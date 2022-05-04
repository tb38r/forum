package web

import (
	"fmt"
	"net/http"

	"forum/dislikes"
	//"strconv"
)

// var CUserIdint int

func (s *myServer) DislikeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("dislike handler running")

		r.ParseForm()

		dislike := r.FormValue("disllike")

		if dislike == "on" {
			dislikes.DislikeButton(s.Db, GuserId, PostIDInt)
		}

		Tpl.ExecuteTemplate(w, "dislikes.html", nil)
	}
}
