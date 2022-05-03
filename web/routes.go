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

		fmt.Println("Len of Requests", len(Requests), Requests)

		if len(Requests) > 0 {
			func() {
				<-a
			}()

			next.ServeHTTP(w, Requests[0])
			fmt.Println(Requests, "\n", "Request sent at ----> ", time.Now())
			Requests = Requests[1:]

		}

		fmt.Println("PART 2 Len of Requests :", len(Requests), Requests)
		fmt.Println()

	}

}

func (s *myServer) Routes(a <-chan time.Time) {
	// http.HandleFunc("/register", srv.LoginAuthHandler())
	http.HandleFunc("/home", Rate(a, HomepageSessionChecker(s.HomepageHandler())))
	http.HandleFunc("/register/", Rate(a, s.RegisterUserHandler()))
	http.HandleFunc("/registerauth", Rate(a, s.RegisterAuthHandler()))
	http.HandleFunc("/login", Rate(a, s.LoginHandler()))
	http.HandleFunc("/loginauth", Rate(a, s.LoginAuthHandler()))
	http.HandleFunc("/logout", Rate(a, s.LogoutHandler()))
	http.HandleFunc("/createpost/", Rate(a, Auth(SessionChecker(s.CreatePostHandler()))))
	http.HandleFunc("/storepost", Rate(a, Auth(SessionChecker(s.StorePostHandler()))))
	http.HandleFunc("/createcomment/", Rate(a, Auth(SessionChecker(s.CreateCommentHandler()))))
	http.HandleFunc("/storecomment", Rate(a, Auth(SessionChecker(s.StoreCommentHandler()))))
	http.HandleFunc("/likes", Rate(a, Auth(SessionChecker(s.LikeHandler()))))
	http.HandleFunc("/showpost/", Rate(a, s.ShowPostHandler()))
	http.HandleFunc("/showcomment/", Rate(a, Auth(SessionChecker(s.ShowCommentHandler()))))
	// http.HandleFunc("/", Rate(a,s.HomepageHandler()))
	// this is for the template css files to run.
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates"))))

}
