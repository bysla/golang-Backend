package routes

import (
	"controller"

	"github.com/gorilla/mux"
)

var MyRouter = mux.NewRouter().StrictSlash(true)

func RouthsP() {

	MyRouter.HandleFunc("/session", controller.Session).Methods("POST")

	MyRouter.HandleFunc("/ongs", controller.InsertOngs).Methods("POST")
	MyRouter.HandleFunc("/ongs", controller.SelectOngs).Methods("GET")

	MyRouter.HandleFunc("/profile", controller.SelectAIncidents).Methods("GET")

	MyRouter.HandleFunc("/incidents", controller.InsertIncidents).Methods("POST")
	MyRouter.HandleFunc("/incidents", controller.SelectIncidents).Methods("GET")
	MyRouter.HandleFunc("/incidents", controller.DeleteAIncidents).Methods("DELETE")

}
