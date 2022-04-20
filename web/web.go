package web

import (
	"forum/web/handlers"
	"forum/web/routes"
	"log"
	"net/http"
)

func OpenServer() {

	// x := server.Server{}
	routes.Routes(handlers.Server{})
	log.Fatal(http.ListenAndServeTLS(":8080", "tls/cert.pem", "tls/key.pem", nil))

}
