package controller

import (
	"dataBase"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

var results []string

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

	fmt.Fprintln(w, string(test1))

}

func DeleteAIncidents(w http.ResponseWriter, r *http.Request) {

	ong_id := r.Header.Get("Authorization")

	ids := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(ids)
	dataBase.DeleteIncidents(id, ong_id)

}

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

	var id int = dataBase.InsertIncidents(title, description, value, ong_id)
	fmt.Fprintf(w, `{"id":"%d"}`, id)
}
