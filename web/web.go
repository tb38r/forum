package web

import (
	"log"
	"net/http"
)

func OpenServer() {

	x := Server{}
	x.Routes()
	log.Fatal(http.ListenAndServeTLS(":8080", "tls/cert.pem", "tls/key.pem", nil))

}
