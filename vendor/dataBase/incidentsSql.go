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

func InsertIncidents(title string, description string, value string, ong_id string) int {
	Statement, _ = Database.Prepare("INSERT INTO incidents (title,description,value,ong_id) VALUES(?,?,?,?)")
	Statement.Exec(title, description, value, ong_id)

	SQL := `SELECT id FROM incidents WHERE id IN (SELECT max(id) FROM incidents)`
	rows, err := Database.Query(SQL)
	if err != nil {
		log.Println(err)
	}
	var id int
	for rows.Next() {
		rows.Scan(&id)
		return id

	}
	return id
}

func SelectIncidents(page int) []DadoIncidents {
	var pages int = ((page - 1) * 5)

	SQL := `SELECT * FROM incidents LIMIT ? OFFSET ?`
	rows, err := Database.Query(SQL, 5, pages)
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

func SelectAIncidents(ongId string) []DadoIncidents {
	SQL := `SELECT * FROM incidents where ong_id = ?`
	rows, err := Database.Query(SQL, ongId)
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

func DeleteIncidents(id int, ong_id string) {
	Statement, _ = Database.Prepare("DELETE FROM incidents WHERE id = ? and ong_id = ?")
	Statement.Exec(id, ong_id)
}
