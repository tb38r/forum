package posts

import (
	"database/sql"
	"fmt"

	"forum/dislikes"
	"forum/likes"
)

type Post struct {
	PostID       int
	UserID       int
	CreationDate string
	PostTitle    string
	PostContent  string
	Image        string
	Edited       bool
}

type HomepagePosts struct {
	PostID       int
	PostTitle    string
	PostUsername string
	CreationDate string
	PostLike     int
	PostDislike  int
}

type ActPage struct {
	UserID       int
	PostTitle    string
	CommentID    int
	CommentText  string
	PostID       int
	CreationDate string
	PostLike     int
}

var db *sql.DB

// type s *web.Server

var LastIns int64

func CreatePosts(db *sql.DB, userID int, title string, content string, image string) {
	stmt, err := db.Prepare("INSERT INTO post (userID, postTitle, postContent, image, creationDate) VALUES (?, ?, ?, ?, strftime('%H:%M %d/%m/%Y','now','localtime'))")
	if err != nil {
		fmt.Println("error preparing statement:", err)
		return
	}
	// defer stmt.Close()
	result, _ := stmt.Exec(userID, title, content, image)

	// checking if the result has been added and the last inserted row
	rowsAff, _ := result.RowsAffected()
	LastIns, _ = result.LastInsertId()
	fmt.Println("rows affected:", rowsAff)
	fmt.Println("last inserted:", LastIns)
}

// Get all the data needed for the hompage
func GetHomepageData(db *sql.DB) []HomepagePosts {
	rows, err := db.Query(`SELECT postID, postTitle, username, creationDate FROM post 
	INNER JOIN users ON users.userID = post.userID;`)
	if err != nil {
		fmt.Println(err)
	}

	postdata := []HomepagePosts{}
	defer rows.Close()
	for rows.Next() {
		var p HomepagePosts
		// fmt.Println(&p.PostID)
		err2 := rows.Scan(&p.PostID, &p.PostTitle, &p.PostUsername, &p.CreationDate)
		p.PostLike = likes.GetPostLikes(db, p.PostID)
		p.PostDislike = dislikes.GetPostDislikes(db, p.PostID)
		postdata = append(postdata, p)
		if err2 != nil {
			fmt.Println(err2)
		}
	}
	return postdata
}

//returns user's comments with their corresponding posts
func ActivityComments(db *sql.DB, userid int) []ActPage {
	rows, err := db.Query(`SELECT post.userID, post.postTitle, comments.commentText, post.postID 
	FROM post, comments
	WHERE post.userID = ?
	AND post.postID = comments.postID
	;`, userid)
	if err != nil {
		fmt.Println(err)
	}
	pac := []ActPage{}
	defer rows.Close()
	for rows.Next() {
		var p ActPage
		err2 := rows.Scan(&p.PostID, &p.PostTitle, &p.CommentText, &p.PostID)
		p.PostLike = likes.GetPostLikes(db, p.PostID)
		pac = append(pac, p)
		if err2 != nil {
			fmt.Println(err2)
		}

	}
	return pac

}

//returns user's liked posts
func ActivityPostLikes(db *sql.DB, userid int) []ActPage {
	rows, err := db.Query(`SELECT DISTINCT likes.userID, likes.postID, post.postTitle
	FROM likes, post
	WHERE likes.userID = ?
	AND likes.postID = post.postID 
	;`, userid)
	if err != nil {
		fmt.Println(err)
	}
	actlikes := []ActPage{}
	defer rows.Close()
	for rows.Next() {
		var p ActPage
		err2 := rows.Scan(&p.UserID, &p.PostID, &p.PostTitle)
		actlikes = append(actlikes, p)
		if err2 != nil {
			fmt.Println(err2)
		}

	}
	return actlikes

}

//returns user's disliked posts
func ActivityPostDislikes(db *sql.DB, userid int) []ActPage {
	rows, err := db.Query(`SELECT DISTINCT dislikes.userID, dislikes.postID, post.postTitle
	FROM dislikes, post
	WHERE dislikes.userID = ?
	AND dislikes.postID = post.postID 
	;`, userid)
	if err != nil {
		fmt.Println(err)
	}
	actpostdislikes := []ActPage{}
	defer rows.Close()
	for rows.Next() {
		var p ActPage
		err2 := rows.Scan(&p.UserID, &p.PostID, &p.PostTitle)
		actpostdislikes = append(actpostdislikes, p)
		if err2 != nil {
			fmt.Println(err2)
		}

	}
	return actpostdislikes

}

//retruns user's liked comments
func ActivityCommentLikes(db *sql.DB, userid int) []ActPage {
	rows, err := db.Query(`SELECT DISTINCT likes.userID, likes.commentID, comments.commentText
	FROM likes, comments
	WHERE likes.userID = 2
	AND likes.commentID = comments.commentID 
	;`, userid)
	if err != nil {
		fmt.Println(err)
	}
	actcommentlikes := []ActPage{}
	defer rows.Close()
	for rows.Next() {
		var p ActPage
		err2 := rows.Scan(&p.UserID, &p.CommentID, &p.CommentText)
		actcommentlikes = append(actcommentlikes, p)
		if err2 != nil {
			fmt.Println(err2)
		}

	}
	return actcommentlikes

}

func FilterHomepageData(db *sql.DB, userID int) []HomepagePosts {
	rows, err := db.Query("SELECT postID, postTitle, username, creationDate FROM post INNER JOIN users ON users.userID =  post.userID WHERE users.userID = ?;", userID)
	if err != nil {
		fmt.Println(err)
	}
	postdata := []HomepagePosts{}
	defer rows.Close()
	for rows.Next() {
		var p HomepagePosts
		err2 := rows.Scan(&p.PostID, &p.PostTitle, &p.PostUsername, &p.CreationDate)
		p.PostLike = likes.GetPostLikes(db, p.PostID)
		p.PostDislike = dislikes.GetPostDislikes(db, p.PostID)
		postdata = append(postdata, p)
		if err2 != nil {
			fmt.Println(err2)
		}
	}
	return postdata
}

// getting the data from one post, and storing the values in the post struct
func GetPostData(db *sql.DB, postID int) []Post {
	rows, err := db.Query("SELECT * FROM post WHERE postID = ?;", postID)
	if err != nil {
		fmt.Println(err)
	}
	postdata := []Post{}
	defer rows.Close()
	for rows.Next() {
		var p Post
		err2 := rows.Scan(&p.PostID, &p.UserID, &p.CreationDate, &p.PostTitle, &p.PostContent, &p.Image, &p.Edited)
		postdata = append(postdata, p)
		if err2 != nil {
			fmt.Println(err2)
		}
	}
	return postdata
}

// selects posts with the same category, based on the category name

func CategoryPagePosts(db *sql.DB, name string) []HomepagePosts {
	rows, err := db.Query(`SELECT post.postID, postTitle, username, creationDate FROM post 
	INNER JOIN category ON category.postID = post.postID 
	INNER JOIN users ON users.userID = post.userID
	WHERE categoryname = ?;`, name)
	if err != nil {
		fmt.Println(err)
	}
	posts := []HomepagePosts{}
	defer rows.Close()
	for rows.Next() {
		var p HomepagePosts
		err2 := rows.Scan(&p.PostID, &p.PostTitle, &p.PostUsername, &p.CreationDate)
		p.PostLike = likes.GetPostLikes(db, p.PostID)
		posts = append(posts, p)
		if err2 != nil {
			fmt.Println(err2)
		}

	}
	return posts
}
