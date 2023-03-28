package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Read(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := Connect()
	
	var read []Person

	err := db.Find(&read).Error

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hasil , err := json.Marshal(read)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(hasil)
}

func ReadById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]
	var person Person

	db := Connect()

	err := db.First(&person, id).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(person)

	w.Write([]byte("Success"))
}

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	db := Connect()
	
		var newPerson Person
		err := json.NewDecoder(r.Body).Decode(&newPerson)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}


		db.Create(&newPerson)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	w.Write([]byte("success"))
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	vars := mux.Vars(r) 
	id := vars["id"]
	var person Person
	
	db := Connect()
	
	db.First(&person, id)

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(&person); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	

	db.Save(&person)

	w.Write([]byte("Update Success"))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]
	
	var person Person

	db := Connect()

	db.First(&person, id)

	db.Delete(&person)

	w.Write([]byte("Delete Success"))
}
