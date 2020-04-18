package routes

import (
	"controller"

	"github.com/gorilla/mux"
)

var MyRouter = mux.NewRouter().StrictSlash(true)

func RouthsP() {

	MyRouter.HandleFunc("/session", controller.SessionRou).Methods("POST")

	MyRouter.HandleFunc("/ongs", controller.SessionRou).Methods("POST")
	MyRouter.HandleFunc("/ongs", controller.SelectOngs).Methods("GET")

	MyRouter.HandleFunc("/profile", controller.InsertOngs).Methods("GET")

	MyRouter.HandleFunc("/incidents", controller.InsertOngs).Methods("POST")
	MyRouter.HandleFunc("/incidents", controller.SelectIncidents).Methods("GET")
	MyRouter.HandleFunc("/incidents", controller.SessionRou).Methods("DELETE")

}
