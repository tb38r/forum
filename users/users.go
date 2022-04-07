package users

import (
	"html/template"
	"net/http"
)

type User struct {
	UserID          int
	Email           string
	Username        string
	Password        string
	Usertype        string
	ExternalLoginId string
}

type AuthUser struct {
	Email        string
	PasswordHash string
}

var tpl = template.Must(template.ParseGlob("templates/*.html"))

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, "register.html")
}
