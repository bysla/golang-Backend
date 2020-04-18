package controller

import (
	"dataBase"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

var results []string

func InsertOngs(w http.ResponseWriter, r *http.Request) {

	ong_id := r.Header.Get("Authorization")
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	title := keyVal["title"]
	description := keyVal["description"]
	value := keyVal["value"]

	fmt.Println(title)
	fmt.Println(description)
	fmt.Println(value)

	dataBase.InsertIncidents(title, description, value, ong_id)
}
