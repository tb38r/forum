package web

import (
	"database/sql"
	"fmt"
	"net/http"

	"forum/likes"
	"forum/posts"
	"forum/users"
)

// eventually need this struct to pass onto the handler with all the data within
// struct fields need to be capitalized, to be used in the templates
type HomepageData struct {
	Username      string
	AllPostTitles []posts.HomepagePosts
	Loggedin      bool
	UserID        int
	Likes         int
	// PostUsername  map[int]string
}

// in chrome this handler is being run twice on localhost:8080, on safari only once (which is what we need) *** UNLESS route is changed from / to /home
func (s *myServer) HomepageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("homepage handler running")
		// checking if user is logged in
		user := users.CurrentUser
		s.Db, _ = sql.Open("sqlite3", "forum.db")
		homepage := posts.GetHomepageData(s.Db)

		PostLikes := likes.GetPostLikes(s.Db, PostIDInt)

		// homePageData := HomepageData{user, postTitles, users.AlreadyLoggedIn(r), GuserId, posts.GetPostUsername(s.Db)}
		homePageData := HomepageData{user, homepage, users.AlreadyLoggedIn(r), GuserId, PostLikes}
		Tpl.ExecuteTemplate(w, "homepage.html", homePageData)
	}
}
