package comments

type Comment struct {
	CommentID    int
	UserID       int
	PostID       int
	CreationDate int
	CommentText  string
	LikesID      int
	DislikesID   int
	Edited       bool
}
