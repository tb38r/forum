package web

import (
	"fmt"
	"net/http"
	"strconv"

	"forum/categories"
	"forum/likes"
	"forum/posts"
)

// type Server server.Server
var UserIdint int

func (s *Server) CreatePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting the user id from the url
		userId := r.URL.Query().Get("userid")
		fmt.Println(userId)
		UserIdint, _ = strconv.Atoi(userId)
		tpl.ExecuteTemplate(w, "createpost.html", nil)
	}
}

func (s *Server) StorePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// s.Db, _ = sql.Open("sqlite3", "forum.db")
		r.ParseForm()

		title := r.FormValue("title")
		content := r.FormValue("content")
		// fmt.Println(UserIdint)
		// adding the post to the database
		posts.CreatePosts(s.Db, UserIdint, title, content)
		// formvalue for buttons. If they have been clicked, the form value returned will be "on"
		manutd := r.FormValue("manutd")
		arsenal := r.FormValue("arsenal")
		chelsea := r.FormValue("chelsea")
		tottenham := r.FormValue("tottenham")
		newcastle := r.FormValue("newcastle")
		mancity := r.FormValue("mancity")
		like := r.FormValue("like")

		// use if statements because we need to enter the cat name instead of the returned value "on"
		if manutd == "on" {
			categories.AddCategory(s.Db, posts.LastIns, "manutd")
		}
		if arsenal == "on" {
			categories.AddCategory(s.Db, posts.LastIns, "arsenal")
		}
		if chelsea == "on" {
			categories.AddCategory(s.Db, posts.LastIns, "chelsea")
		}
		if newcastle == "on" {
			categories.AddCategory(s.Db, posts.LastIns, "newcastle")
		}
		if tottenham == "on" {
			categories.AddCategory(s.Db, posts.LastIns, "tottenham")
		}
		if mancity == "on" {
			categories.AddCategory(s.Db, posts.LastIns, "mancity")
		}
		if like == "on" {
			likes.LikeButton(s.Db, GuserId, 1)
		}
		fmt.Println("title:", title, "content:", content)

		tpl.ExecuteTemplate(w, "storepost.html", "Post stored!")
	}
}
