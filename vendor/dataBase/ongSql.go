package dataBase

import (
	"log"
	"util"
)

type DadoOngs struct {
	Id       string
	Name     string
	Email    string
	Whatsapp string
	City     string
	Uf       string
}

func InsertOngs(name string, email string, whatsapp string, city string, uf string) string {
	id, _ := util.RandomHex(4)
	Statement, _ = Database.Prepare("INSERT INTO ongs (id,name,email,whatsapp,city,uf) VALUES(?,?,?,?,?,?)")
	Statement.Exec(id, name, email, whatsapp, city, uf)

	return id
}

func SelectOngs() []DadoOngs {
	SQL := `SELECT * FROM ongs`
	rows, err := Database.Query(SQL)
	if err != nil {
		log.Println(err)
	}

	var id, name, email, whatsapp, city, uf string
	var listaOngs []DadoOngs

	for rows.Next() {
		rows.Scan(&id, &name, &email, &whatsapp, &city, &uf)
		listaOngs = append(listaOngs, DadoOngs{
			Id:       id,
			Name:     name,
			Email:    email,
			Whatsapp: whatsapp,
			City:     city,
			Uf:       uf,
		})
	}
	return listaOngs

}

func SelectAOngs(ongId string) string {
	SQL := `SELECT * FROM ongs WHERE id = ? LIMIT 1 `
	rows, err := Database.Query(SQL, ongId)
	if err != nil {
		log.Println(err)
	}

	var id, name, email, whatsapp, city, uf string
	var nome string

	for rows.Next() {
		rows.Scan(&id, &name, &email, &whatsapp, &city, &uf)
		nome = name
	}
	return nome

}
