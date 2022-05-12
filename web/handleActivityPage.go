package web

import (
	"database/sql"
	"fmt"
	"forum/posts"
	"net/http"
)

type ActivityPage struct {
	Posts    []posts.HomepagePosts
	Likes    []posts.Post
	Dislikes []posts.Post
	Comments []posts.Post
}

func (s *myServer) ActivityPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Db, _ = sql.Open("sqlite3", "forum.db")

		var data ActivityPage

		userFilter := posts.FilterHomepageData(s.Db, GuserId)
		fmt.Println(userFilter)

		data.Posts = userFilter

		Tpl.ExecuteTemplate(w, "activitypage.html", data)

	}
}
