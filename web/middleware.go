package web

import (
	"fmt"
	"forum/users"
	"net/http"
)

func Auth(HandlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		c, err := r.Cookie(users.CurrentUser)
		fmt.Println(c.Value)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// _, user := users.DbSessions[users.CurrentUser]
		// if !user {
		// 	fmt.Println("Checking loop ------------------->", users.CurrentUser)
		// if !users.SessionExists(users.CurrentUser) {

		// }
		HandlerFunc.ServeHTTP(w, r)
	}
}
