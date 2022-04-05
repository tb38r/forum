package data

import "database/sql"

const dbName = "forum.db"

func CreateDB() *sql.DB {
	db, _ := sql.Open("sqlite3", dbName)
	db.Exec("create table if not exists testTable (id integer primary key,username text, surname text,age integer,university text)")

	return db
}
