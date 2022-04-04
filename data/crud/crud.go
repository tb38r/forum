package crud

import (
	"database/sql"
	"strconv"
)

type User struct {
	id         int
	username   string
	surname    string
	age        int
	university string
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}

	// catch to error.

}

func AddUser(db *sql.DB, username string, surname string, age int, university string) {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into testTable (username,surname,age,university) values (?,?,?,?)")
	_, err := stmt.Exec(username, surname, age, university)
	checkError(err)
	tx.Commit()
}

func GetUser(db *sql.DB, id2 int) User {
	rows, err := db.Query("select * from testTable")
	checkError(err)
	for rows.Next() {
		var tempUser User
		err = rows.Scan(&tempUser.id, &tempUser.username, &tempUser.surname, &tempUser.age, &tempUser.university)
		checkError(err)
		if tempUser.id == id2 {
			return tempUser
		}
	}
	return User{}
}

func UpdateUser(db *sql.DB, id2 int, username string, surname string, age int, university string) {
	sage := strconv.Itoa(age)
	sid := strconv.Itoa(id2)

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("update testTable set username=?, surname=?, age=?, university=? where id=?")
	_, err := stmt.Exec(username, surname, sage, university, sid)
	checkError(err)
	tx.Commit()
}

func DeleteUser(db *sql.DB, id2 int) {
	sid := strconv.Itoa(id2) // int to string
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("delete from testTable where id=?")
	_, err := stmt.Exec(sid)
	checkError(err)
	tx.Commit()
}
