package posts

import (
	"database/sql"
	"fmt"
	"html/template"
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

// this global variable is needed in Handleposts to ge the postid from the last inserted row.
var LastIns int64
var tpl = template.Must(template.ParseGlob("templates/*.html"))

func CreatePosts(db *sql.DB, userID int, title string, content string) {
	stmt, err := db.Prepare("INSERT INTO post (userID, postTitle, postContent, creationDate) VALUES (?, ?, ?, datetime('now', 'localtime'))")
	if err != nil {
		fmt.Println("error preparing statement:", err)
		return
	}
	// defer stmt.Close()
	result, _ := stmt.Exec(userID, title, content)

	// db.Close()

	// checking if the result has been added and the last inserted row
	rowsAff, _ := result.RowsAffected()
	LastIns, _ = result.LastInsertId()
	fmt.Println("rows affected:", rowsAff)
	fmt.Println("last inserted:", LastIns)
}

// func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
// 	// getting the user id from the url
// 	userId := r.URL.Query().Get("userid")
// 	UserIdint, _ = strconv.Atoi(userId)
// 	tpl.ExecuteTemplate(w, "createpost.html", nil)
// }

// func StorePostHandler(w http.ResponseWriter, r *http.Request) {
// 	db, _ = sql.Open("sqlite3", "forum.db")
// 	r.ParseForm()

// 	title := r.FormValue("title")
// 	content := r.FormValue("content")
// 	// fmt.Println(UserIdint)
// 	// adding the post to the database
// 	CreatePosts(db, UserIdint, title, content)

// 	fmt.Println("title:", title, "content:", content)

// 	tpl.ExecuteTemplate(w, "storepost.html", "Post stored!")
// }
