package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", Read).Methods("GET")
	r.HandleFunc("/users/{id}", ReadById).Methods("GET")
	r.HandleFunc("/users/create", Create).Methods("POST")
	r.HandleFunc("/users/update/{id}", Update).Methods("PUT")
	r.HandleFunc("/users/delete/{id}", Delete).Methods("DELETE")

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", r)
}