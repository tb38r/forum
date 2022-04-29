package web

import (
	"log"
	"net/http"
	"time"
)

func OpenServer(a <-chan time.Time) {

	x := Server{}
	x.Routes(a)
	log.Fatal(http.ListenAndServeTLS(":8080", "tls/cert.pem", "tls/key.pem", nil))

}
