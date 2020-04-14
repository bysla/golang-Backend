package controller

import (
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
