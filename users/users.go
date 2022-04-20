package users

import (
	"database/sql"
	"fmt"
	"html/template"
	"net"
	"net/http"

	//"net/mail"
	"strings"
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

//var tpl *template.Template
var currentUser string

//var dbUsers = map[string]User{}
var DbSessions = make(map[string]string)

var tpl = template.Must(template.ParseGlob("templates/*.html"))

//dbUsers["yonas@hotmail.com"] = User{Username: "yonas123", Email: "yonas@hotmail.com"}

var db *sql.DB

//nil

// this func registers a users username, password(as a hash) and email
func RegisterUser(db *sql.DB, username string, hash []byte, email string) {
	// db, _ = sql.Open("sqlite3", "forum.db")
	stmt, err := db.Prepare("INSERT INTO users (username, hash, email) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println("error preparing statement:", err)
		return
	}
	// defer stmt.Close()

	result, _ := stmt.Exec(username, hash, email)
	db.Close()

	// checking if the result has been added and the last inserted row
	rowsAff, _ := result.RowsAffected()
	lastIns, _ := result.LastInsertId()
	fmt.Println("rows affected:", rowsAff)
	fmt.Println("last inserted:", lastIns)
}

func ValidEmail(email string) bool {
	i := strings.Index(email, "@")
	fmt.Println("i:", i)

	domain := email[i+1:]
	fmt.Println("Domain: ", domain)

	_, err := net.LookupMX(domain)
	// _, err2 := mail.ParseAddress(email)
	if err != nil {
		// fmt.Println("invalid email")
		return false
	}

	return true
}

func AlreadyLoggedIn(r *http.Request) bool {

	c, err := r.Cookie(currentUser)
	if err != nil {
		return false
	}
	_, ok := DbSessions[c.Name]
	return ok
}

func SessionExists(s string) bool {

	if _, user := DbSessions[s]; user {
		return true
	}
	return false
}
