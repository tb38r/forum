package web

import (
	"fmt"
	"net/http"
)

// eventually need this struct to pass onto the handler with all the data within
// type HomepageData struct {
// 	// allPostTitles []string
// 	Loggedin?  bool

// }

// in chrome this handler is being run twice on localhost:8080, on safari only once (which is what we need) *** UNLESS route is changed from / to /home
func (s *Server) HomepageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("homepage handler running")
		Tpl.ExecuteTemplate(w, "homepage.html", nil)
	}
}
