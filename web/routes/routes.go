package routes

import (
	"forum/web/server"
	"net/http"
)

func Routes(srv server.Server) {
	http.HandleFunc("/register/", srv.RegisterUserHandler())
	http.HandleFunc("/registerauth", srv.RegisterAuthHandler())
	http.HandleFunc("/login", srv.LoginHandler())
	http.HandleFunc("/loginauth", srv.LoginAuthHandler())
	http.HandleFunc("/logout", srv.LogoutHandler())
	http.HandleFunc("/createpost/", srv.CreatePostHandler())
	http.HandleFunc("/storepost", srv.StorePostHandler())

}
