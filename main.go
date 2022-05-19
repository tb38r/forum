package main

import (
	"forum/database"
	"forum/web"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type RateLimiter struct {
	limiter <-chan time.Time
	seconds time.Duration
}

func main() {
	//tt
	rate := RateLimiter{}

	rate.seconds = 2
	rate.limiter = time.Tick(rate.seconds * time.Second)

	database.CreateDB()
	web.OpenServer(rate.limiter)
}
