package main

import (
	"database/sql"
	"forum/database"
	"forum/web"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		log.Fatal(err)
	}
	database.CreateDB(db)

	web.OpenServer()

}
