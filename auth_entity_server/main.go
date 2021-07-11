package main

import (
	"encoding/json"
	"net/http"
)

type Accounts struct {
	Accounts []Account `json:"accounts"`
}

type Account struct {
	Name string `json:"name"`
}

func accountsHandler(w http.ResponseWriter, r *http.Request) {
	// TODO : implement authentication entity logic
	a := Accounts{[]Account{Account{"John"}, Account{"Sary"}}}

	j, err := json.Marshal(a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(j)
}

func main() {
	http.HandleFunc("/", accountsHandler)
	http.ListenAndServe(":3000", nil)
}
