package main

import (
	"log"
	"net/http"
	"routes"
)

func main() {
	//dataBase.CreateDB()
	routes.RouthsP()
	log.Fatal(http.ListenAndServe(":80", routes.MyRouter))

}
