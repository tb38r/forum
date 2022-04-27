package web

import (
	"fmt"
	"net/http"
)

func (s *Server) HomepageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("homepage handler running")
		Tpl.ExecuteTemplate(w, "homepage.html", nil)
	}
}
