package web

import (
	"forum/posts"
	"net/http"
)

type ActivityPage struct {
	Posts    []posts.HomepagePosts
	CWP      []posts.ActPage //commentswithposts
	APL      []posts.ActPage //activity post likes
	APD      []posts.ActPage //activity post dislikes
	Comments []posts.Post
}

func (s *myServer) ActivityPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var data ActivityPage

		data.Posts = posts.FilterHomepageData(s.Db, GuserId)

		data.CWP = posts.ActivityComments(s.Db, GuserId)

		//value := posts.ActivityPostLikes(s.Db, GuserId)
		// for _, item := range value {
		// 	fmt.Print("UserID: ", item.UserID, "\t")
		// 	fmt.Print("Title: ", item.PostTitle, "\t")
		// 	fmt.Print("Comment: ", item.CommentText, "\t")
		// 	fmt.Println()

		// }

		data.APL = posts.ActivityPostLikes(s.Db, GuserId)

		data.APD = posts.ActivityPostDislikes(s.Db, GuserId)

		Tpl.ExecuteTemplate(w, "activitypage.html", data)

	}
}
