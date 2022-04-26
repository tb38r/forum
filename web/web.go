package web

import (
	"forum/posts"
	"forum/users"

	"log"
	"net/http"
)

func OpenServer() {

	// x := server.Server{}
	UserRoutes(users.Server{})
	PostRoutes(posts.Server{})
	log.Fatal(http.ListenAndServeTLS(":8080", "tls/cert.pem", "tls/key.pem", nil))

}
