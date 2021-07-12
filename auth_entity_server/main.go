package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Accounts struct {
	Accounts []Account `json:"accounts"`
}

type Account struct {
	Name string `json:"name"`
}

type AuthDelegateRequest struct {
	BindingMessage  string `json:"binding_message"`
	LoginHint       string `json:"login_hint"`
	ConsentRequired bool   `json:"is_consent_required"`
	Scope           string `json:"scope"`
	AcrValues       string `json:"acr_values"`
}

func accountsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incoming request from %s.", r.Host)
	// TODO : implement authentication entity logic

	// receive json and Authorization header for bearer token for token authentication afterwards
	var authDelegateRequest AuthDelegateRequest
	json.NewDecoder(r.Body).Decode(&authDelegateRequest)
	log.Printf("    acr_values : %s.", authDelegateRequest.AcrValues)
	log.Printf("    binding_message : %s.", authDelegateRequest.BindingMessage)
	//log.Printf("    is_consent_required : %s.", authDelegateRequest.ConsentRequired)
	log.Printf("    login_hint : %s.", authDelegateRequest.LoginHint)
	log.Printf("    scope : %s.", authDelegateRequest.Scope)

}

func main() {
	log.Println("Auth Entity Server booted.")
	http.HandleFunc("/", accountsHandler)
	http.ListenAndServe(":3001", nil)
}
