package dataBase

import (
	"database/sql"
	"util"

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

	//database, _ := sql.Open("sqlite3","./data/data.db")
}

func InsertOngs(name string, email string, whatsapp string, city string, uf string) {
	id, _ := util.RandomHex(4)
	Statement, _ = Database.Prepare("INSERT INTO ongs (id,name,email,whatsapp,city,uf) VALUES(?,?,?,?,?,?)")
	Statement.Exec(id, name, email, whatsapp, city, uf)
}
