package web

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"
)

func OpenServer(a <-chan time.Time) {

	mainserver := myServer{
		serve: &http.Server{
			Addr:              ":8080",
			ReadHeaderTimeout: 0,
			Handler:           nil,
			TLSConfig: &tls.Config{
				MinVersion: tls.VersionTLS13,
				PreferServerCipherSuites: true,
			},
		},
	}

	mainserver.Routes(a)

	// x.serve.ReadHeaderTimeout = 2
	// x.serve.Handler =

	// srv := &http.Server{
	// 	Addr:              ":8080",
	// 	ReadHeaderTimeout: 10,
	// 	Handler:           nil,
	// 	TLSConfig: &tls.Config{
	// 		MinVersion: tls.VersionTLS10,
	// 	},
	// }
	//srv.ReadHeaderTimeout = 10

	log.Fatal(mainserver.serve.ListenAndServeTLS("tls/cert.pem", "tls/key.pem"))

}
