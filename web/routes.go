package web

import (
	"fmt"
	"net/http"
	"time"
)

var Requests []*http.Request

func Rate(a <-chan time.Time, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		Requests = append(Requests, r)

		if len(Requests) > 0 {
			func() {
				<-a
			}()

			next.ServeHTTP(w, Requests[0])
			fmt.Println(Requests, "\n", "Request sent at ----> ", time.Now())
			Requests = Requests[1:]

		}

		fmt.Println()

	}

}

func (s *myServer) Routes(a <-chan time.Time) {
	// http.HandleFunc("/register", srv.LoginAuthHandler())
	http.HandleFunc("/register/", Rate(a, s.RegisterUserHandler()))
	http.HandleFunc("/registerauth", Rate(a, s.RegisterAuthHandler()))
	http.HandleFunc("/login", Rate(a, s.LoginHandler()))
	http.HandleFunc("/loginauth", Rate(a, s.LoginAuthHandler()))
	http.HandleFunc("/logout", Rate(a, s.LogoutHandler()))
	http.HandleFunc("/createpost/", Rate(a, Auth(SessionChecker(s.CreatePostHandler()))))
	http.HandleFunc("/storepost", Rate(a, Auth(SessionChecker(s.StorePostHandler()))))
	http.HandleFunc("/", Rate(a, s.HomepageHandler()))

}
