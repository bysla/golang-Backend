package routes

import (
	"controller"

	"github.com/gorilla/mux"
)

var MyRouter = mux.NewRouter().StrictSlash(true)

func RouthsP() {

	MyRouter.HandleFunc("/session", controller.SessionRou).Methods("POST")

	MyRouter.HandleFunc("/ongs", controller.SessionRou).Methods("POST")
	MyRouter.HandleFunc("/ongs", controller.SessionRou).Methods("GET")

	MyRouter.HandleFunc("/profile", controller.SessionRou).Methods("GET")

	MyRouter.HandleFunc("/incidents", controller.SessionRou).Methods("POST")
	MyRouter.HandleFunc("/incidents", controller.SessionRou).Methods("GET")
	MyRouter.HandleFunc("/incidents", controller.SessionRou).Methods("DELETE")

}
