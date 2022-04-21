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

		// use a switch case instead because we need to enter the cat name instead of the returned value "on"

		switch "on" {
		case manutd:
			{
				categories.AddCategory(users.Db, LastIns, "manutd")
			}
		case arsenal:
			{
				categories.AddCategory(users.Db, LastIns, "arsenal")
			}
		case chelsea:
			{
				categories.AddCategory(users.Db, LastIns, "chelsea")
			}
		case tottenham:
			{
				categories.AddCategory(users.Db, LastIns, "tottenham")
			}
		case mancity:
			{
				categories.AddCategory(users.Db, LastIns, "mancity")
			}
		case newcastle:
			{
				categories.AddCategory(users.Db, LastIns, "newcastle")
			}
		}

		// fmt.Println("last inserted check", LastIns)

		// postIdQuery := "SELECT postID FROM post WHERE userID = ?"

		// fmt.Println("this is the user id from store post handler", UserIdint)

		// row := users.Db.QueryRow(postIdQuery, 1)
		// var postID int
		// err := row.Scan(&postID)
		// fmt.Println("lets see if it picks up post id", postID)
		// if err != sql.ErrNoRows {
		// 	fmt.Println("postID or UserID does not exist, err:", err)
		// 	server.Tpl.ExecuteTemplate(w, "storepost.html", "Post not stored!")
		// 	return
		// }

		fmt.Println("title:", title, "content:", content)

		server.Tpl.ExecuteTemplate(w, "storepost.html", "Post stored!")
	}
}
