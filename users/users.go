package users

import (
	"fmt"
	"html/template"
	"net/http"
	"unicode"
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

func registerAuthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("********registerAuthHandler running*******")

	r.ParseForm()

	username := r.FormValue("username")

	var nameAlphaNumeric = true

	//check whether each character in username AlphaNumeric

	for _, char := range username {
		if unicode.IsLetter(char) == false && unicode.IsNumber(char) == false {
			nameAlphaNumeric = false
		}
	}

	//check username length is 6 < x < 10

	var usernameLength bool

	if len(username) <= 10 && len(username) >= 6 {
		usernameLength = true
	}

	//check password is valid given all conditions

	password := r.FormValue("password")
	fmt.Println("password", password, "\npswdLength",len(password))

	var pswdLowercase, pswdUpperCase, pswdNumber, pswdSpecical, pswdNoSpaces, passwordLength bool

	pswdNoSpaces = true

	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			pswdLowercase = true
		case unicode.IsUpper(char):
			pswdUpperCase = true
		case unicode.IsNumber(char):
			pswdNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			pswdSpecical = true
		case unicode.IsSpace(int32(char)):
			pswdNoSpaces = false
		}
	}
	//check password length is 5 < x < 20
	if len(password) <= 20 && len(password) >= 5{
		passwordLength = true
	}


}