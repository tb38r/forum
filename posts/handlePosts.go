package posts

import (
	"database/sql"
	"fmt"

	"forum/users"
	"forum/web/server"
	"net/http"
	"strconv"

	uuid "github.com/satori/go.uuid"
)

type Server server.Server

var req http.Request
var wr http.ResponseWriter

func (s *Server) CreatePostHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		cookiecheck(w, r)

		// getting the user id from the url
		userId := r.URL.Query().Get("userid")
		UserIdint, _ = strconv.Atoi(userId)
		server.Tpl.ExecuteTemplate(w, "createpost.html", nil)

	}

}

func (s *Server) StorePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Db, _ = sql.Open("sqlite3", "forum.db")
		r.ParseForm()

		title := r.FormValue("title")
		content := r.FormValue("content")
		// fmt.Println(UserIdint)
		// adding the post to the database
		CreatePosts(s.Db, UserIdint, title, content)

		fmt.Println("title:", title, "content:", content)

		server.Tpl.ExecuteTemplate(w, "storepost.html", "Post stored!")
	}
}

var cookiecheck = func(w http.ResponseWriter, r *http.Request) {
	if users.SessionExists(users.CurrentUser) {
		fmt.Println()
		fmt.Println("Session Exists/ID", users.DbSessions)
		fmt.Println()

		// Read the cookie,if there are any
		cookie, err := r.Cookie(users.CurrentUser)

		//i.e session exists (on map) but no active cookie on (presumably) new client
		if err != nil {

			// create new cookie for user on this client
			id := uuid.Must(uuid.NewV4())
			c := &http.Cookie{
				Name:  users.CurrentUser,
				Value: id.String(),
			}

			http.SetCookie(w, c)
			users.DbSessions[users.CurrentUser] = c.Value
			server.Tpl.ExecuteTemplate(w, "loginauth.html", "session created after reassigning ID in map")
			fmt.Println("Map Values reassigned for new client log in: ", users.DbSessions)
			fmt.Println()

			return

			//session exists but differing UUID,  logout/close session
		} else if cookie.Value != users.DbSessions[users.CurrentUser] {

			//expire cookie as there's an active session elsewhere
			for _, cookie := range r.Cookies() {
				fmt.Println("-------Test Delete 3")

				if cookie.Name == users.CurrentUser {
					fmt.Println("checking if loop entered--------------------------------------------------------")
					fmt.Println("1st max age check ----> ", cookie.MaxAge, cookie.Value)

					cookie.MaxAge = -1
					//cookie.Value = ""
					//cookie.Expires = time.Date(2022, time.April, 21, 0000, 00, 00, 00, time.UTC)
					fmt.Println("2nd max age check ----> ", cookie.MaxAge, cookie.Value)
				}
				http.SetCookie(w, cookie)
			}

			//redirects to log in page when prior session exists
			http.Redirect(w, r, "/logout", http.StatusSeeOther)

			fmt.Println("Cookie deleted due to pre-existing session")
			fmt.Println()
			return

			//UUID matches that within map, active session, no conflicts
		} else if cookie.Value == users.DbSessions[users.CurrentUser] {
			server.Tpl.ExecuteTemplate(w, "loginauth.html", "Active session no changes")
			fmt.Println("Active session on client, no changes made")
			fmt.Println()
			return

		}

	}

}

func (s *Server) LogoutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, _ := r.Cookie(users.CurrentUser)

		// delete the session
		delete(users.DbSessions, c.Name)

		// remove the cookie
		c = &http.Cookie{
			Name:   users.CurrentUser,
			Value:  "",
			MaxAge: -1,
		}
		http.SetCookie(w, c)

		fmt.Println("User logged out and redirected to the log-in page")

		http.Redirect(w, r, "/login", http.StatusSeeOther)

	}
}
