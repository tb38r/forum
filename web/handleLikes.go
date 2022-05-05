package web

import (
	"fmt"
	"net/http"
	"strconv"

	"forum/likes"
)

// var CUserIdint int

func (s *myServer) LikeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("like handler running")

		likes.LikeButton(s.Db, GuserId, PostIDInt)

		fmt.Println(PostIDInt)

		SPostID := strconv.Itoa(PostIDInt)

		http.Redirect(w, r, "/showpost/?postid="+SPostID, http.StatusSeeOther)
	}
}
