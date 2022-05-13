package web

import (
	"net/http"

	"forum/posts"
)

type ActivityPage struct {
	Posts    []posts.HomepagePosts
	CWP      []posts.ActPage // commentswithposts
	Dislikes []posts.Post
	Comments []posts.Post
}

func (s *myServer) ActivityPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data ActivityPage

		userFilter := posts.FilterHomepageData(s.Db, GuserId)

		data.Posts = userFilter

		data.CWP = posts.ActivityComments(s.Db, GuserId)

		Tpl.ExecuteTemplate(w, "activitypage.html", data)
	}
}