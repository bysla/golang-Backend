package controller

import (
	"dataBase"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Session(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	id := keyVal["id"]

	test1 := dataBase.SelectAOngs(id)

	fmt.Fprintf(w, `{"name":"%s"}`, test1)
}
