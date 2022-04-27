package web

import (
	"fmt"
	"forum/comments"
	"net/http"
	"strconv"
)

var CUserIdint int

func (s *Server) CreateCommentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		userID := r.URL.Query().Get("userid")
		CUserIdint, _ = strconv.Atoi(userID)
		Tpl.ExecuteTemplate(w, "createcomment.html", nil)

	}
}

func (s *Server) StoreCommentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()
		contentComment := r.FormValue("content")

		comments.CreateComment(s.Db, GuserId, contentComment)

		fmt.Println("content: ", contentComment)
		Tpl.ExecuteTemplate(w, "storecomment.html", nil)

	}
}
