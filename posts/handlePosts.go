package posts

import (
	"database/sql"
	"fmt"

	"forum/users"
	"forum/web/server"
	"net/http"
	"strconv"
)

type Server server.Server

func (s *Server) CreatePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting the user id from the url
		userId := r.URL.Query().Get("userid")
		UserIdint, _ = strconv.Atoi(userId)
		server.Tpl.ExecuteTemplate(w, "createpost.html", nil)
	}
}

func (s *Server) StorePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users.Db, _ = sql.Open("sqlite3", "forum.db")
		r.ParseForm()

		title := r.FormValue("title")
		content := r.FormValue("content")

		// formvalue for buttons. If they have been clicked, the form value returned will be "on"
		manutd := r.FormValue("manutd")
		arsenal := r.FormValue("arsenal")
		chelsea := r.FormValue("chelsea")
		tottenham := r.FormValue("tottenham")
		liverpool := r.FormValue("liverpool")
		mancity := r.FormValue("mancity")

		// use a switch case instead because we need to enter the cat name instead of the returned value "on"

		switch "on" {
		case manutd:
			{
				fmt.Println("Man utd was clicked")
			}
		case arsenal:
			{

			}
		case chelsea:
			{

			}
		case tottenham:
			{

			}
		case mancity:
			{

			}
		case liverpool:
			{

			}
		}
		// adding the post to the database
		CreatePosts(users.Db, UserIdint, title, content)

		fmt.Println("title:", title, "content:", content)

		server.Tpl.ExecuteTemplate(w, "storepost.html", "Post stored!")
	}
}
