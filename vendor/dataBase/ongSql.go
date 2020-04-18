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

func InsertOngs(name string, email string, whatsapp string, city string, uf string) {
	id, _ := util.RandomHex(4)
	Statement, _ = Database.Prepare("INSERT INTO ongs (id,name,email,whatsapp,city,uf) VALUES(?,?,?,?,?,?)")
	Statement.Exec(id, name, email, whatsapp, city, uf)
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
