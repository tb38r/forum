package web

import (
	"database/sql"
	"net/http"

	"forum/posts"
	"forum/users"
)

// eventually need this struct to pass onto the handler with all the data within
// struct fields need to be capitalized, to be used in the templates
type HomepageData struct {
	Username string
	AllPosts []posts.HomepagePosts
	Loggedin bool
	UserID   int
	// PostsByCategory []posts.CategoryPagePosts
	// PostUsername  map[int]string
}

// in chrome this handler is being run twice on localhost:8080, on safari only once (which is what we need) *** UNLESS route is changed from / to /home
func (s *myServer) HomepageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := users.CurrentUser
		s.Db, _ = sql.Open("sqlite3", "forum.db")
		homepage := posts.GetHomepageData(s.Db)
		category := r.FormValue("category")
		homePageData := HomepageData{user, homepage, users.AlreadyLoggedIn(r), GuserId}
		homePageFilter := r.FormValue("userfilter")
		// Choosing which data is passed into the homepage based on the filter chosen
		if len(category) < 1 && len(homePageFilter) < 1 {
			Tpl.ExecuteTemplate(w, "homepage.html", homePageData)
		} else if len(category) > 0 {
			categoryFilter := posts.CategoryPagePosts(s.Db, category)
			homePageData = HomepageData{user, categoryFilter, users.AlreadyLoggedIn(r), GuserId}
			Tpl.ExecuteTemplate(w, "homepage.html", homePageData)
		} else if len(homePageFilter) > 0 {
			userFilter := posts.FilterHomepageData(s.Db, GuserId)
			homePageData = HomepageData{user, userFilter, users.AlreadyLoggedIn(r), GuserId}
			Tpl.ExecuteTemplate(w, "homepage.html", homePageData)
		}
	}
}
