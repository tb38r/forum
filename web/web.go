package web

import (
	"forum/posts"
	"forum/users"
	"forum/web/routes"
	"log"
	"net/http"
)

func OpenServer() {

	// x := server.Server{}
	routes.UserRoutes(users.Server{})
	routes.PostRoutes(posts.Server{})
	log.Fatal(http.ListenAndServeTLS(":8080", "tls/cert.pem", "tls/key.pem", nil))

}
