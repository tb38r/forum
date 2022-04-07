package web

import (
	"fmt"
	"forum/users"
	"log"
	"net/http"
)

func OpenServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World Secure!")
	})

	http.HandleFunc("/register/", users.RegisterUser)
	http.HandleFunc("/registerauth", users.RegisterAuthHandler)
	log.Fatal(http.ListenAndServeTLS(":8080", "tls/cert.pem", "tls/key.pem", nil))
}
