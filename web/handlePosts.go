package web

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"forum/categories"
	"forum/posts"
)

type PostPageData struct {
	PostId   int
	Title    string
	Content  string
	Comments []string
	Likes    int
	Dislikes int
}

// type Server server.Server
var UserIdint int
var PostIDInt int

func (s *myServer) CreatePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting the user id from the url
		userId := r.URL.Query().Get("userid")
		UserIdint, _ = strconv.Atoi(userId)
		Tpl.ExecuteTemplate(w, "createpost.html", nil)
	}
}

func (s *myServer) StorePostHandler() http.HandlerFunc {
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
		fmt.Println("title:", title, "content:", content)

		Tpl.ExecuteTemplate(w, "storepost.html", "Post stored!")
	}
}

func (s *myServer) ShowPostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// had to open the database here as it wasnt picking the correct post everytime without this.
		s.Db, _ = sql.Open("sqlite3", "forum.db")
		// get the postId and display the post and its contents
		postID := r.URL.Query().Get("postid")
		PostIDInt, _ = strconv.Atoi(postID)

		// r.ParseForm()

		// like := r.FormValue("like")

		Tpl.ExecuteTemplate(w, "showpost.html", posts.GetPostData(s.Db, PostIDInt))
	}
}
