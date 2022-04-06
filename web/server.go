package web

import (
	"fmt"
	"log"
	"net/http"
)

func OpenServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World Secure!")
	})

	log.Fatal(http.ListenAndServeTLS(":8080", "localhost.crt", "localhost.key", nil))
}
