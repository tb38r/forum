package web

import (
	"net/http"
)

func (s *Server) Routes() {
	// http.HandleFunc("/register", srv.LoginAuthHandler())
	http.HandleFunc("/register/", s.RegisterUserHandler())
	http.HandleFunc("/registerauth", s.RegisterAuthHandler())
	http.HandleFunc("/login", s.LoginHandler())
	http.HandleFunc("/loginauth", s.LoginAuthHandler())
	http.HandleFunc("/logout", s.LogoutHandler())
	http.HandleFunc("/createpost/", Auth(SessionChecker(s.CreatePostHandler())))
	http.HandleFunc("/storepost", Auth(SessionChecker(s.StorePostHandler())))
	http.HandleFunc("/", s.HomepageHandler())

}
