package web

import (
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
	CommentNote       []Notify
	LikeNote          []Notify
	DisLikeNote       []Notify
}

func (s *myServer) ActivityPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var data ActivityPage

		data.Notification = (len(CommentNotify(s.Db)) + len(LikesNotify(s.Db)) + len(DisLikesNotify(s.Db)))

		if data.Notification > 0 {
			data.Nbool = true
		}

		data.CommentNote = CommentNotify(s.Db)

		data.LikeNote = LikesNotify(s.Db)

		data.DisLikeNote = DisLikesNotify(s.Db)

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

		Tpl.ExecuteTemplate(w, "activitypage.html", data)

		func() {
			ResetCommentNotified(s.Db)
		}()

		func() {
			ResetLikesNotified(s.Db)
		}()

		func() {
			ResetDisLikesNotified(s.Db)
		}()

	}
}
