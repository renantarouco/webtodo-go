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
	todoRouter.HandleFunc("/", RetrieveItemsHandler).Methods(http.MethodGet)
	todoRouter.HandleFunc("/new", CreateItemHandler).Methods(http.MethodGet, http.MethodPost)
	todoRouter.HandleFunc("/{id}", UpdateItemHandler).Methods(http.MethodGet, http.MethodPost)
	todoRouter.HandleFunc("/delete/{id}", DeleteItemHandler).Methods(http.MethodGet)
	todoRouter.HandleFunc("/notfound", NotFoundHandler).Methods(http.MethodGet)
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println(err.Error())
	}
}
