package routes

import (
	"forum/web/handlers"
	"net/http"
)

func Routes(srv handlers.Server) {
	//http.HandleFunc("/register/", srv.handlers.RegisterUserHandler())
	http.HandleFunc("/register/", srv.RegisterUserHandler())
	http.HandleFunc("/registerauth", srv.RegisterAuthHandler())
	http.HandleFunc("/login", srv.LoginHandler())
	http.HandleFunc("/loginauth", srv.LoginAuthHandler())
	http.HandleFunc("/logout", srv.LogoutHandler())
	http.HandleFunc("/createpost/", srv.CreatePostHandler())
	http.HandleFunc("/storepost", srv.StorePostHandler())

}
