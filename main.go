package main

import "dataBase"

func main() {
	//routes.RouthsP()
	//log.Fatal(http.ListenAndServe(":80", routes.MyRouter))
	dataBase.CreateDB()
	dataBase.InsertOngs("luis", "luisfelipe@gmail.com", "5519999610009", "sjbv", "sp")

}
