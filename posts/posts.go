package posts

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
)

type Post struct {
	PostID       int
	UserID       int
	CommentID    int
	CategoryID   int
	CreationDate int
	PostTitle    string
	PostText     string
	PostImage    string
	LikesID      int
	DislikesID   int
	Edited       bool
}

var db *sql.DB

var tpl = template.Must(template.ParseGlob("templates/*.html"))

func (*Post) createPosts(db *sql.DB, userID int, title string, content string) {
	stmt, err := db.Prepare("INSERT INTO post (userID, postTitle, postContent) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println("error preparing statement:", err)
		return
	}
	// defer stmt.Close()

	result, _ := stmt.Exec(userID, title, content)
	db.Close()

	// checking if the result has been added and the last inserted row
	rowsAff, _ := result.RowsAffected()
	lastIns, _ := result.LastInsertId()
	fmt.Println("rows affected:", rowsAff)
	fmt.Println("last inserted:", lastIns)
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "createpost.html", nil)
}

func StorePostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	title := r.FormValue("title")
	content := r.FormValue("content")

	fmt.Println("title:", title, "content:", content)

	tpl.ExecuteTemplate(w, "storepost.html", nil)
}
