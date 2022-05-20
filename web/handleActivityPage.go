package web

import (
	"fmt"
	"forum/posts"
	"forum/users"
	"net/http"
	"strconv"
)

type ActivityPage struct {
	Posts             []posts.HomepagePosts
	CommentsWithPosts []posts.ActPage
	LikedPosts        []posts.ActPage
	DislikedPosts     []posts.ActPage
	LikedComments     []posts.ActPage
	DislikedComments  []posts.ActPage
	Comments          []posts.Post
	Username          string
	LoggedIn          bool
	UserID            int
	Nbool             bool
	Notification      int
	CommentNote       []CommNotify
}

func (s *myServer) ActivityPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data ActivityPage

		if len(CommentNotify(s.Db)) > 0 {
			data.Nbool = true
		}

		data.Notification = len(CommentNotify(s.Db))

		data.CommentNote = CommentNotify(s.Db)
		fmt.Println(data.CommentNote)

		data.Posts = posts.UsersPostsHomepageData(s.Db, GuserId)

		data.CommentsWithPosts = posts.ActivityComments(s.Db, GuserId)

		data.LikedPosts = posts.ActivityPostLikes(s.Db, GuserId)

		data.DislikedPosts = posts.ActivityPostDislikes(s.Db, GuserId)

		data.LikedComments = posts.ActivityCommentLikes(s.Db, GuserId)

		data.DislikedComments = posts.ActivityCommentDislikes(s.Db, GuserId)
		data.Username = users.CurrentUser
		data.LoggedIn = users.AlreadyLoggedIn(r)
		data.UserID = GuserId
		SuserID := strconv.Itoa(GuserId)

		if string(r.URL.RawQuery[len(r.URL.RawQuery)-1]) != SuserID {
			http.Error(w, "Incorrect user request made!", http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// go func() {
		// 	ResetCommentNotified(s.Db)
		// }()

		Tpl.ExecuteTemplate(w, "activitypage.html", data)

		go func(){
			fmt.Println("-----ENTERS GO FUNC")
			ResetCommentNotified(s.Db)
		}()
	}
}
