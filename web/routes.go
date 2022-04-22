package web

import (
	"forum/posts"
	"forum/users"

	"forum/web/server"
	"net/http"
)

type rtr server.Server

func UserRoutes(srv users.Server) {
	http.HandleFunc("/register", srv.LoginAuthHandler())
	//http.HandleFunc("/register/", srv.handlers.RegisterUserHandler())
	http.HandleFunc("/register/", srv.RegisterUserHandler())
	http.HandleFunc("/registerauth", srv.RegisterAuthHandler())
	http.HandleFunc("/login", srv.LoginHandler())
	http.HandleFunc("/loginauth", srv.LoginAuthHandler())
	http.HandleFunc("/logout", srv.LogoutHandler())

}

func PostRoutes(srv posts.Server) {
	http.HandleFunc("/createpost/", Auth(srv.CreatePostHandler()))
	http.HandleFunc("/storepost", srv.StorePostHandler())
}
