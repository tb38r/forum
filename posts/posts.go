package posts

type Post struct {
	PostID       int
	UserID       int
	CommentID    int
	CategoryID   int
	CreationDate int
	PostText     string
	PostImage    string
	LikesID      int
	DislikesID   int
	Edited       bool
}
