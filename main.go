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
	//dataBase.DeleteIncidents(33, "94370077")
	routes.RouthsP()
	log.Fatal(http.ListenAndServe(":80", routes.MyRouter))
	//dataBase.InsertIncidents("ajuda", "preciso de money", "150", "a6474b90")

}
