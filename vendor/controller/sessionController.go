package controller

import (
	"dataBase"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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

	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}
	n, _ := strconv.Atoi(page)
	test1, err := json.Marshal(dataBase.SelectIncidents(n))
	if err != nil {
		return
	}

	fmt.Fprintln(w, string(test1))

}

func SelectAIncidents(w http.ResponseWriter, r *http.Request) {

	ong_id := r.Header.Get("Authorization")
	test1, err := json.Marshal(dataBase.SelectAIncidents(ong_id))
	if err != nil {
		return
	}

	fmt.Println(string(test1)) //ele printa como [{},{},{},....{}]

	fmt.Fprintln(w, string(test1))

}

func DeleteAIncidents(w http.ResponseWriter, r *http.Request) {

	ong_id := r.Header.Get("Authorization")

	ids := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(ids)
	dataBase.DeleteIncidents(id, ong_id)

}

var results []string

func InsertIncidents(w http.ResponseWriter, r *http.Request) {

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

	var id int = dataBase.InsertIncidents(title, description, value, ong_id)
	fmt.Fprintf(w, `{"id":"%d"}`, id)
}

func InsertOngs(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	name := keyVal["name"]
	email := keyVal["email"]
	whatsapp := keyVal["whatsapp"]
	city := keyVal["city"]
	uf := keyVal["uf"]

	var id string = dataBase.InsertOngs(name, email, whatsapp, city, uf)
	fmt.Fprintf(w, `{"id":"%s"}`, id)
}
