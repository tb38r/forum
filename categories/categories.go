package categories

import (
	"database/sql"
	"fmt"
)

type Category struct {
	CategoryName string
	CategoryID   int
	PostID       int
}

var db *sql.DB

func AddCategory(db *sql.DB, postID int64, categoryname string) {
	stmt, err := db.Prepare("INSERT INTO category (postID, categoryname) VALUES (?, ?)")

	if err != nil {
		fmt.Println("error preparing statment", err)
		return
	}

	result, _ := stmt.Exec(postID, categoryname)

	// checking if the result has been added and the last inserted row
	rowsAff, _ := result.RowsAffected()
	lastIns, _ := result.LastInsertId()
	fmt.Println("rows affected:", rowsAff)
	fmt.Println("last inserted:", lastIns)

}
