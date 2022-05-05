package web

import (
	"fmt"
	"net/http"
	"strconv"

	"forum/dislikes"
)

// var CUserIdint int

func (s *myServer) DislikeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("dislike handler running")

		dislikes.DislikeButton(s.Db, GuserId, PostIDInt)

		fmt.Println(PostIDInt)

		SPostID := strconv.Itoa(PostIDInt)

		http.Redirect(w, r, "/showpost/?postid="+SPostID, http.StatusSeeOther)
	}
}
