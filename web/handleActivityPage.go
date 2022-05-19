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
}

func (s *myServer) ActivityPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data ActivityPage

		data.Posts = posts.UsersPostsHomepageData(s.Db, GuserId)

		data.CommentsWithPosts = posts.ActivityComments(s.Db, GuserId)

		//value := posts.ActivityCommentLikes(s.Db, GuserId)
		// for _, item := range value {
		// 	fmt.Print("UserID: ", item.UserID, "\t")
		// 	fmt.Print("Title: ", item.PostTitle, "\t")
		// 	fmt.Print("Comment: ", item.CommentText, "\t")
		// 	fmt.Println()

		// }

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
			// fmt.Fprintln(w, r.URL.Path)
			// fmt.Fprintln(w, r.URL.RawQuery)
			return
		}

		// if r.URL.Path != "/activitypage/"+r.URL.RawQuery {
		// 	//	http.Error(w, "Incorrect user request made!", http.StatusBadRequest)
		// 	fmt.Fprintln(w, r.URL.Path)
		// 	fmt.Fprintln(w, r.URL.RawQuery)
		// 	return
		// }
		Tpl.ExecuteTemplate(w, "activitypage.html", data)
	}
}
