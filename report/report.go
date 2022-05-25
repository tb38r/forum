package report

import (
	"database/sql"
	"fmt"
)

type Report struct {
	ReportID int
	Mod      string
	PostID   int
}

var LastIns int64

func ReportButton(db *sql.DB, username string, reporttype string, postID int) {
	stmt, err := db.Prepare("INSERT INTO report(username, reporttype, postID) VALUES(?, ?, ?)")
	if err != nil {
		fmt.Println("error preparing statement:", err)
		return
	}

	result, _ := stmt.Exec(username, reporttype, postID)
	rowsAff, _ := result.RowsAffected()
	LastIns, _ = result.LastInsertId()
	fmt.Println("rows affected:", rowsAff)
	fmt.Println("last inserted", LastIns)
}

func GetReportType(db *sql.DB, reportID int) string {
	var reportType string
	err := db.QueryRow("SELECT reporttype FROM report WHERE reportID= ?;", reportID).Scan(&reportType)
	if err != nil {
		fmt.Println("error from get report", err)
	}
	return reportType
}
func GetMod(db *sql.DB, reportID int) string {
	var username string
	err := db.QueryRow("SELECT username FROM users WHERE reportID= ?;", reportID).Scan(&username)
	if err != nil {
		fmt.Println("error from get mod", err)
	}
	return username
}
