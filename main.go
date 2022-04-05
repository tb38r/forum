package main

import (
	"devto/data"
	"devto/data/crud"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const dbName = "forum.db"

func main() {
	db := data.CreateDB()
	crud.AddUser(db, "LLvzi omur", "tekin", 24, "Sakarya University") // added data to database
	crud.AddUser(db, "Ken", "Thompson", 75, "California university")  //update data to database
	crud.DeleteUser(db, 1)
	//fmt.Println(crud.GetUser(db, 2)) // printing the user

	crud.UpdateUser(db, 2, "Ken", "Thompson", 45, "01 founders")

	fmt.Println(crud.GetUser(db, 2)) // printing the user
}
