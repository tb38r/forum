package web

import (
	"database/sql"
	"net/http"
	"text/template"
)

var Tpl = template.Must(template.ParseGlob("templates/*.html"))

type Server struct {
	Db     *sql.DB
	Router *http.ServeMux
}

// func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	s.ServeHTTP(w, r)
// }
