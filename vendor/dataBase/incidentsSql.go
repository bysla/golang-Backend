package dataBase

import (
	"log"
)

type DadoIncidents struct {
	Id          int
	Title       string
	Description string
	Value       string
	Ong_id      string
}

func InsertIncidents(title string, description string, value string, ong_id string) {
	Statement, _ = Database.Prepare("INSERT INTO incidents (title,description,value,ong_id) VALUES(?,?,?,?)")
	Statement.Exec(title, description, value, ong_id)
}

func SelectIncidents() []DadoIncidents {
	SQL := `SELECT * FROM incidents`
	rows, err := Database.Query(SQL)
	if err != nil {
		log.Println(err)
	}

	var id int
	var title, description, value, ong_id string
	var listaIncidents []DadoIncidents

	for rows.Next() {
		rows.Scan(&id, &title, &description, &value, &ong_id)
		listaIncidents = append(listaIncidents, DadoIncidents{
			Id:          id,
			Title:       title,
			Description: description,
			Value:       value,
			Ong_id:      ong_id,
		})
	}
	return listaIncidents

}
