package web

import (
	"net/http"
)

func (s *Server) Routes() {
	// http.HandleFunc("/register", srv.LoginAuthHandler())
	http.HandleFunc("/home", s.HomepageHandler())
	http.HandleFunc("/register/", s.RegisterUserHandler())
	http.HandleFunc("/registerauth", s.RegisterAuthHandler())
	http.HandleFunc("/login", s.LoginHandler())
	http.HandleFunc("/loginauth", s.LoginAuthHandler())
	http.HandleFunc("/logout", s.LogoutHandler())
	http.HandleFunc("/createpost/", Auth(SessionChecker(s.CreatePostHandler())))
	http.HandleFunc("/storepost", Auth(SessionChecker(s.StorePostHandler())))
	// this is for the template css files to run.
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates"))))

	http.HandleFunc("/likes", s.LikeHandler())
}
