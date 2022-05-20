package web

import (
	"database/sql"
	"fmt"
	"net/http"

	"forum/posts"
	"forum/users"
)

// eventually need this struct to pass onto the handler with all the data within
// struct fields need to be capitalized, to be used in the template
type HomepageData struct {
	Username     string
	AllPosts     []posts.HomepagePosts
	LoggedIn     bool
	UserID       int
	Nbool        bool
	Notification int
	// PostUsername  map[int]string
}

// in chrome this handler is being run twice on localhost:8080, on safari only once (which is what we need) *** UNLESS route is changed from / to /home
func (s *myServer) HomepageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		user := users.CurrentUser
		s.Db, _ = sql.Open("sqlite3", "forum.db")

		fmt.Println("COMMENT USERNAME MAP", CommentUsername(s.Db))

		homepage := posts.GetHomepageData(s.Db)
		var x bool

		if len(CommentUsername(s.Db)) > 0 {
			x = true
		}
		notify := len(CommentUsername(s.Db))

		homePageData := HomepageData{user, homepage, users.AlreadyLoggedIn(r), GuserId, x, notify}
		category := r.FormValue("category")
		homePageFilter := r.FormValue("userfilter")

		// Choosing which data is passed into the homepage based on the filter chosen
		if len(category) < 1 && len(homePageFilter) < 1 {
			Tpl.ExecuteTemplate(w, "homepage.html", homePageData)
		} else if len(category) > 0 {
			categoryFilter := posts.CategoryPagePosts(s.Db, category)
			homePageData = HomepageData{user, categoryFilter, users.AlreadyLoggedIn(r), GuserId, x, notify}
			Tpl.ExecuteTemplate(w, "homepage.html", homePageData)
		} else if homePageFilter == "Created Post" {
			userFilter := posts.UsersPostsHomepageData(s.Db, GuserId)
			homePageData = HomepageData{user, userFilter, users.AlreadyLoggedIn(r), GuserId, x, notify}
			Tpl.ExecuteTemplate(w, "homepage.html", homePageData)
		} else if homePageFilter == "Liked Posts" {
			userFilter := posts.UsersLikesHomepageData(s.Db, GuserId)
			homePageData = HomepageData{user, userFilter, users.AlreadyLoggedIn(r), GuserId, x, notify}
			Tpl.ExecuteTemplate(w, "homepage.html", homePageData)
		}
	}
}
