package web

import (
	"forum/comments"
	"net/http"
	"strconv"
)

func (s *Server) CreateCommentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("userid")
		comments.UserIDint, _ = strconv.Atoi(userID)
		Tpl.ExecuteTemplate(w, "createcomment.html", nil)

	}
}

func (s *Server) StoreCommentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()
		content := r.FormValue("content")

		comments.CreateComment(s.Db, comments.UserIDint, content)

		Tpl.ExecuteTemplate(w, "storecomment.html", nil)

	}
}
