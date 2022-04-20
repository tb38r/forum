package server

import (
	"net/http"
	"text/template"
)

var Tpl = template.Must(template.ParseGlob("templates/*.html"))

type Server struct {
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.ServeHTTP(w, r)
}
