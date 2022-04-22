package web

import (
	"forum/users"
	"net/http"
)

func Auth(HandlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		_, user := users.DbSessions[users.CurrentUser]
		if !user {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		HandlerFunc.ServeHTTP(w, r)
	}
}
