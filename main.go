package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	todoRouter := router.PathPrefix("/todos").Subrouter().StrictSlash(true)
	// Todas as rotas mapeadas e com métodos permitidos filtrados.
	todoRouter.HandleFunc("/", RetrieveItems).Methods(http.MethodGet)
	todoRouter.HandleFunc("/new", CreateItem).Methods(http.MethodGet, http.MethodPost)
	todoRouter.HandleFunc("/{id}", UpdateItem).Methods(http.MethodGet, http.MethodPut)
	todoRouter.HandleFunc("/{id}", DeleteItem).Methods(http.MethodDelete)
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println(err.Error())
	}
}
