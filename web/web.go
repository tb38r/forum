package web

import (
	"forum/web/routes"
	"forum/web/server"
	"log"
	"net/http"
)

func OpenServer() {

	x := server.Server{}
	routes.Routes(x)
	log.Fatal(http.ListenAndServeTLS(":8080", "tls/cert.pem", "tls/key.pem", nil))

	// server := http.Server{
	// 	Addr: ":8080",
	// }

	// if err := server.ListenAndServeTLS("tls/cert.pem", "tls/key.pem"); err != nil {
	// 	panic(err)
	// }

}
