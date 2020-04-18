package controller

import (
	"dataBase"
	"encoding/json"
	"fmt"
	"net/http"
)

func SessionRou(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Fprintf(w, "POST")
	} else {
		http.Error(w, "Invalid request method.", 405)
	}
}

func SelectOngs(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		test1, err := json.Marshal(dataBase.SelectOngs())
		if err != nil {
			return
		}

		fmt.Println(string(test1)) //ele printa como [{},{},{},....{}]

		fmt.Fprintln(w, string(test1))
	} else {
		http.Error(w, "Invalid request method.", 405)
	}
}

func SelectIncidents(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		test1, err := json.Marshal(dataBase.SelectIncidents())
		if err != nil {
			return
		}

		fmt.Println(string(test1)) //ele printa como [{},{},{},....{}]

		fmt.Fprintln(w, string(test1))
	} else {
		http.Error(w, "Invalid request method.", 405)
	}
}

func InsertOngs(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	r.ParseForm()

}
