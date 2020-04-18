package main

import (
	"log"
	"net/http"
	"routes"
)

func main() {
	//dataBase.CreateDB()
	//var bt []dataBase.DadoOngs = dataBase.SelectOngs()
	//fmt.Printf("%v\n", bt)
	routes.RouthsP()
	log.Fatal(http.ListenAndServe(":80", routes.MyRouter))
	//dataBase.CreateDB()
	//dataBase.InsertIncidents("ajuda", "preciso de money", "150", "a6474b90")

}
