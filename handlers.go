package main

import "net/http"

import "fmt"

// Implementação das funcionalidades CRUD básicas.
// Create e Update são rotas que terão funcionalidades diferentes para cada
// método HTTP descrito.

func RetrieveItems(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "listing to dos")
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintln(w, "showing to-do form page")
	case http.MethodPost:
		fmt.Fprintln(w, "inserting to-do")
	}
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintln(w, "showing to-do form page")
	case http.MethodPut:
		fmt.Fprintln(w, "updating to-do")
	}
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "deleting to do")
}
