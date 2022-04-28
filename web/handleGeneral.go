package web

import (
	"fmt"
	"forum/users"
	"net/http"
)

// eventually need this struct to pass onto the handler with all the data within
// struct fields need to be capitalized, to be used in the templates
type HomepageData struct {
	Username string
	// AllPostTitles []string
	Loggedin bool
}

// in chrome this handler is being run twice on localhost:8080, on safari only once (which is what we need) *** UNLESS route is changed from / to /home
func (s *Server) HomepageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("homepage handler running")
		// checking if user is logged in
		user := users.CurrentUser
		homePageData := HomepageData{user, users.AlreadyLoggedIn(r)}
		Tpl.ExecuteTemplate(w, "homepage.html", homePageData)
	}
}
