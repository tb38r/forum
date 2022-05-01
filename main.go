package main

import (
	"forum/database"
	"forum/tls"
	"forum/web"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type RateLimiter struct {
	limiter <-chan time.Time
	seconds time.Duration
}

func main() {


	rate := RateLimiter{}

	rate.seconds = 1
	rate.limiter = time.Tick(rate.seconds * time.Second)

	database.CreateDB()

	tls.Pemcert()
	tls.Pemkey()

	web.OpenServer(rate.limiter)
}
