package controller

import (
	"dataBase"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SelectOngs(w http.ResponseWriter, r *http.Request) {

	test1, err := json.Marshal(dataBase.SelectOngs())
	if err != nil {
		return
	}

	fmt.Fprintln(w, string(test1))

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
