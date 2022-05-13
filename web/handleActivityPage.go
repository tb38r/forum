package web

import (
	"fmt"
	"forum/posts"
	"net/http"
)

type ActivityPage struct {
	Posts             []posts.HomepagePosts
	CommentsWithPosts []posts.ActPage
	LikedPosts        []posts.ActPage
	DislikedPosts     []posts.ActPage
	LikedComments     []posts.ActPage
	Comments          []posts.Post
}

func (s *myServer) ActivityPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var data ActivityPage

		data.Posts = posts.FilterHomepageData(s.Db, GuserId)

		data.CommentsWithPosts = posts.ActivityComments(s.Db, GuserId)

		value := posts.ActivityCommentLikes(s.Db, GuserId)
		for _, item := range value {
			fmt.Print("UserID: ", item.UserID, "\t")
			fmt.Print("Title: ", item.PostTitle, "\t")
			fmt.Print("Comment: ", item.CommentText, "\t")
			fmt.Println()

		}

		data.LikedPosts = posts.ActivityPostLikes(s.Db, GuserId)

		data.DislikedPosts = posts.ActivityPostDislikes(s.Db, GuserId)

		data.LikedComments = posts.ActivityCommentLikes(s.Db, GuserId)

		Tpl.ExecuteTemplate(w, "activitypage.html", data)

	}
}
