package main

import (
	"forum/database"
	"forum/web"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database.CreateDB()
	web.OpenServer()

	// for _, email := range []string{
	// 	"good@exmaple.com",
	// 	"arnoldmutungi@live.co.uk",
	// 	"sarmad@yonas.com",
	// } {
	// 	fmt.Printf("%v valid: %t\n", email, users.ValidEmail(email))
	// }
}
