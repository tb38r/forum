package handlers

import (
	"database/sql"
	"fmt"
	"forum/posts"
	"forum/users"
	"forum/web/server"
	"net/http"
	"strconv"
)

func (s *Server) CreatePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting the user id from the url
		userId := r.URL.Query().Get("userid")
		posts.UserIdint, _ = strconv.Atoi(userId)
		server.Tpl.ExecuteTemplate(w, "createpost.html", nil)
	}
}

func (s *Server) StorePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users.Db, _ = sql.Open("sqlite3", "forum.db")
		r.ParseForm()

		title := r.FormValue("title")
		content := r.FormValue("content")
		// fmt.Println(UserIdint)
		// adding the post to the database
		posts.CreatePosts(users.Db, posts.UserIdint, title, content)

		fmt.Println("title:", title, "content:", content)

		server.Tpl.ExecuteTemplate(w, "storepost.html", "Post stored!")
	}
}
