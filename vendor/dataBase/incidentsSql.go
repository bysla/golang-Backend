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

func SelectAIncidents(ongId string) []DadoIncidents {
	SQL := `SELECT * FROM incidents where ong_id = ` + ongId
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

func DeleteIncidents(id int, ong_id string) {
	//SQL := `DELETE FROM incidents WHERE id = '` + string(id) + `' and ong_id = '` + ong_id + `'`
	Statement, _ = Database.Prepare("DELETE FROM incidents WHERE id =? and ong_id = ?")
	Statement.Exec(id, ong_id)

}
