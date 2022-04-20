package main

import (
	"forum/database"
	"forum/web"
	_ "forum/web/server"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database.CreateDB()
	web.OpenServer()
}
