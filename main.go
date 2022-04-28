package main

import (
	"forum/database"
	"forum/web"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var (
	limiter <-chan time.Time
	seconds time.Duration
)

func main() {
	seconds = 2
	limiter = time.Tick(seconds * time.Second)

	database.CreateDB()
	web.OpenServer(limiter)
}
