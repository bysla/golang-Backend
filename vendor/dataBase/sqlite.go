package dataBase

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var Database, _ = sql.Open("sqlite3", "./data/db.sqlite")
var Statement, _ = Database.Prepare("")

func CreateDB() {

	Statement, _ := Database.Prepare("CREATE TABLE IF NOT EXISTS ongs" +
		"(id INTEGER PRIMARY KEY," +
		" name TEXT NOT NULL," +
		" email TEXT NOT NULL," +
		" whatsapp TEXT NOT NULL," +
		" city TEXT NOT NULL," +
		" uf TEXT NOT NULL)")

	Statement.Exec()

	Statement, _ = Database.Prepare("CREATE TABLE IF NOT EXISTS incidents" +
		"(id INTEGER PRIMARY KEY AUTO_INCREMENT," +
		" title TEXT NOT NULL," +
		" description TEXT NOT NULL," +
		" value TEXT NOT NULL," +
		" ong_id TEXT NOT NULL)")

	Statement.Exec()

	//database, _ := sql.Open("sqlite3","./data/data.db")
}
