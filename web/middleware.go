package web

import (
	"fmt"
	"forum/users"
	"net/http"
)

func Auth(HandlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		c, err := r.Cookie(users.CurrentUser)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		fmt.Println(c.Value)

		HandlerFunc.ServeHTTP(w, r)
	}
}
