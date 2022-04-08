package users

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"unicode"

	"golang.org/x/crypto/bcrypt"
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

var db *sql.DB

// this func registers a users username, password(as a hash)
func registerUser(db *sql.DB, username string, hash []byte) {
	// db, _ = sql.Open("sqlite3", "forum.db")
	stmt, err := db.Prepare("INSERT INTO users (username, hash) VALUES (?, ?)")

	if err != nil {
		fmt.Println("error preparing statement:", err)
		return
	}
	// defer stmt.Close()

	result, _ := stmt.Exec(username, hash)
	db.Close()

	// checking if the result has been added and the last inserted row
	rowsAff, _ := result.RowsAffected()
	lastIns, _ := result.LastInsertId()
	fmt.Println("rows affected:", rowsAff)
	fmt.Println("last inserted:", lastIns)

}

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "register.html", nil)
}

func RegisterAuthHandler(w http.ResponseWriter, r *http.Request) {
	db, _ = sql.Open("sqlite3", "forum.db")
	fmt.Println("********registerAuthHandler running*******")

	r.ParseForm()

	username := r.FormValue("username")

	nameAlphaNumeric := true

	// check whether each character in username AlphaNumeric

	for _, char := range username {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			nameAlphaNumeric = false
		}
	}

	// check username length is 6 < x < 10

	var usernameLength bool

	if len(username) <= 10 && len(username) >= 6 {
		usernameLength = true
	}

	// check password is valid given all conditions

	password := r.FormValue("password")
	fmt.Println("password", password, "\npswdLength", len(password))

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
	// check password length is 5 < x < 20
	if len(password) <= 20 && len(password) >= 5 {
		passwordLength = true
	}

	if !nameAlphaNumeric || !usernameLength || !pswdLowercase || !pswdUpperCase || !pswdNumber || !pswdSpecical || !pswdNoSpaces || !passwordLength {
		tpl.ExecuteTemplate(w, "register.html", "Please check your username or password")
		return
	}

	// check if username already exists
	stmt := "SELECT userID FROM users WHERE username = ?"
	row := db.QueryRow(stmt, username)
	var uID string
	err := row.Scan(&uID)
	if err != sql.ErrNoRows {
		fmt.Println("username already exists, err:", err)
		tpl.ExecuteTemplate(w, "register.html", "username already exists")
		return
	}

	//  create hash from password
	var hash []byte

	hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("bcrypt err:", err)
		tpl.ExecuteTemplate(w, "register.html", "there was a problem registering user")
		return
	}

	fmt.Println("hash:", hash)
	fmt.Println("string(hash)", string(hash))
	registerUser(db, username, hash)
	fmt.Fprintf(w, "congrats your account has been successfully created")

}
