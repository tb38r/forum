package posts

import (
	"database/sql"
	"fmt"

	"forum/categories"
	"forum/users"
	"forum/web/server"
	"net/http"
	"strconv"
)

type Server server.Server

// this global variable for the userId will be used to get the id from create post handler (in url), and passed onto
// the storepost handler to add as the foreign key of the posts table
var UserIdint int
var userID string

func (s *Server) CreatePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting the user id from the url
		userID = r.URL.Query().Get("userid")
		UserIdint, _ = strconv.Atoi(userID)
		server.Tpl.ExecuteTemplate(w, "createpost.html", nil)
	}
}

func (s *Server) StorePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users.Db, _ = sql.Open("sqlite3", "forum.db")
		r.ParseForm()

		title := r.FormValue("title")
		content := r.FormValue("content")

		// adding the post to the database
		CreatePosts(users.Db, UserIdint, title, content)

		// formvalue for buttons. If they have been clicked, the form value returned will be "on"
		manutd := r.FormValue("manutd")
		arsenal := r.FormValue("arsenal")
		chelsea := r.FormValue("chelsea")
		tottenham := r.FormValue("tottenham")
		newcastle := r.FormValue("newcastle")
		mancity := r.FormValue("mancity")

		// use if statements because we need to enter the cat name instead of the returned value "on"
		if manutd == "on" {
			categories.AddCategory(users.Db, LastIns, "manutd")
		}
		if arsenal == "on" {
			categories.AddCategory(users.Db, LastIns, "arsenal")
		}
		if chelsea == "on" {
			categories.AddCategory(users.Db, LastIns, "chelsea")
		}
		if newcastle == "on" {
			categories.AddCategory(users.Db, LastIns, "newcastle")
		}
		if tottenham == "on" {
			categories.AddCategory(users.Db, LastIns, "tottenham")
		}
		if mancity == "on" {
			categories.AddCategory(users.Db, LastIns, "mancity")
		}

		fmt.Println("title:", title, "content:", content)

		server.Tpl.ExecuteTemplate(w, "storepost.html", "Post stored!")
	}
}
