package server

import (
	"database/sql"
	"fmt"
	"forum/posts"
	"forum/users"
	"net/http"
	"strconv"
	"text/template"
	"unicode"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var tpl = template.Must(template.ParseGlob("templates/*.html"))

type Server struct {
	db          *sql.DB
	currentUser string
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.ServeHTTP(w, r)
}

func (s *Server) RegisterUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "register.html", nil)
	}
}

func (s *Server) RegisterAuthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.db, _ = sql.Open("sqlite3", "forum.db")
		fmt.Println("********registerAuthHandler running*******")

		r.ParseForm()

		username := r.FormValue("username")
		fmt.Println("username", username)

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

		email := r.FormValue("email")
		fmt.Println("email", email)

		if !users.ValidEmail(email) {
			tpl.ExecuteTemplate(w, "register.html", "Please enter a valid email address")
			return
		}

		// check if email already exists
		emailStmt := "SELECT userID FROM users WHERE email = ?"
		rowE := s.db.QueryRow(emailStmt, email)
		var uID string
		err := rowE.Scan(&uID)
		if err != sql.ErrNoRows {
			fmt.Println("email already exists, err:", err)
			tpl.ExecuteTemplate(w, "register.html", "email already exists")
			return
		}

		// check if username already exists
		userStmt := "SELECT userID FROM users WHERE username = ?"
		rowU := s.db.QueryRow(userStmt, username)
		var uIDs string
		error := rowU.Scan(&uIDs)
		if error != sql.ErrNoRows {
			fmt.Println("username already exists, err:", error)
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
		users.RegisterUser(s.db, username, hash, email)
		fmt.Fprintf(w, "congrats your account has been successfully created")

	}
}

func (s *Server) LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("login handler running")
		fmt.Println("checking bool-----> ", users.AlreadyLoggedIn(r))
		tpl.ExecuteTemplate(w, "login.html", nil)
	}
}

func (s *Server) LoginAuthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// we need to figure out whether we have to close the database at some point to save resources.
		s.db, _ = sql.Open("sqlite3", "forum.db")
		fmt.Println("login authHandler running")
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")

		s.currentUser = username
		fmt.Println("username:", username, "password", password)
		// get password(hash form) from db to compare with users supplied password
		var hash string
		stmt := "SELECT hash FROM users WHERE username = ?"
		row := s.db.QueryRow(stmt, username)
		err := row.Scan(&hash)
		fmt.Println("hash from db:", hash)
		if err != nil {
			fmt.Println("error with username, may not exist")
			// keep the message to the user a little more vague, so hackers dont know whether you entered an incorrect username or password
			tpl.ExecuteTemplate(w, "login.html", "check username and password")
			return
		}

		// get userId to pass onto createpost handler

		var userID int

		stmt2 := "SELECT userID FROM users WHERE username = ?"
		row2 := s.db.QueryRow(stmt2, username)
		err2 := row2.Scan(&userID)
		fmt.Println("userID from db:", userID)
		if err2 != nil {
			fmt.Println("user not found in db")
		}

		err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
		if err == nil {
			if users.SessionExists(username) {
				fmt.Println()
				fmt.Println("Session Exists/ID", users.DbSessions)
				fmt.Println()

				// Read the cookie,if there are any
				cookie, err := r.Cookie(username)

				//i.e session exists (on map) but no active cookie on (presumably) new client
				if err != nil {

					// create new cookie for user on this client
					id := uuid.Must(uuid.NewV4())
					c := &http.Cookie{
						Name:  username,
						Value: id.String(),
					}

					http.SetCookie(w, c)
					users.DbSessions[username] = c.Value
					tpl.ExecuteTemplate(w, "loginauth.html", "session created after reassigning ID in map")
					fmt.Println("Map Values reassigned for new client log in: ", users.DbSessions)
					fmt.Println()

					return

					//session exists but differing UUID,  logout/close session
				} else if cookie.Value != users.DbSessions[username] {

					//expire cookie as there's an active session elsewhere
					for _, cookie := range r.Cookies() {
						fmt.Println("-------Test Delete 3")

						if cookie.Name == username {

							cookie.MaxAge = -1

							http.SetCookie(w, cookie)
						}
					}

					//redirects to log in page when prior session exists
					http.Redirect(w, r, "/login", http.StatusSeeOther)

					fmt.Println("Cookie deleted due to pre-existing session")
					fmt.Println()
					return

					//UUID matches that within map, active session, no conflicts
				} else if cookie.Value == users.DbSessions[username] {
					tpl.ExecuteTemplate(w, "loginauth.html", "Active session no changes")
					fmt.Println("Active session on client, no changes made")
					fmt.Println()
					return

				}

			} else {

				//NO ACTIVE SESSION/FIRST TIME
				id := uuid.Must(uuid.NewV4())
				c := &http.Cookie{
					Name:  username,
					Value: id.String(),
				}

				http.SetCookie(w, c)
				users.DbSessions[username] = c.Value
				tpl.ExecuteTemplate(w, "loginauth.html", userID)

				/////////remove///////////////////
				fmt.Println("sessionbool", users.SessionExists(username))

				for _, cookie := range r.Cookies() {
					fmt.Println()
					fmt.Println("Name : ", cookie.Name)
					fmt.Println("Value/UUID : ", cookie.Value)
				}

				fmt.Println("First time log-in successful")
				fmt.Println()
				/////////////////////////////////////
				return

			}

		}

		fmt.Println("incorrect password")
		tpl.ExecuteTemplate(w, "login.html", "check username and password")
	}
}

func (s *Server) LogoutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, _ := r.Cookie(s.currentUser)

		// delete the session
		delete(users.DbSessions, c.Name)

		// remove the cookie
		c = &http.Cookie{
			Name:   s.currentUser,
			Value:  "",
			MaxAge: -1,
		}
		http.SetCookie(w, c)

		fmt.Println("User logged out and redirected to the log-in page")

		http.Redirect(w, r, "/login", http.StatusSeeOther)

	}
}

func (s *Server) CreatePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting the user id from the url
		userId := r.URL.Query().Get("userid")
		posts.UserIdint, _ = strconv.Atoi(userId)
		tpl.ExecuteTemplate(w, "createpost.html", nil)
	}
}

func (s *Server) StorePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.db, _ = sql.Open("sqlite3", "forum.db")
		r.ParseForm()

		title := r.FormValue("title")
		content := r.FormValue("content")
		// fmt.Println(UserIdint)
		// adding the post to the database
		posts.CreatePosts(s.db, posts.UserIdint, title, content)

		fmt.Println("title:", title, "content:", content)

		tpl.ExecuteTemplate(w, "storepost.html", "Post stored!")
	}
}
