package report

import (
	"database/sql"
	"fmt"
)

type Report struct {
	ReportID int
	UserID   int
	PostID   int
}

var LastIns int64

func ReportButton(db *sql.DB, userID int, postID int) {
	stmt, err := db.Prepare("INSERT INTO report(userID, postID) VALUES(?, ?)")
	if err != nil {
		fmt.Println("error preparing statement:", err)
		return
	}

	result, _ := stmt.Exec(userID, postID)
	rowsAff, _ := result.RowsAffected()
	LastIns, _ = result.LastInsertId()
	fmt.Println("rows affected:", rowsAff)
	fmt.Println("last inserted", LastIns)
}
