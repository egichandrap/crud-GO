package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"rest-go-demo/databases"
	"rest-go-demo/entity"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllPerson(w http.ResponseWriter, r *http.Request) {
	var persons []entity.Person
	databases.Connector.Find(&persons)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(persons)
}

func GetPersonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var person entity.Person
	databases.Connector.First(&person, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var person entity.Person
	json.Unmarshal(requestBody, &person)
	databases.Connector.Save(&person)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(person)
}

func DeletPersonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var person entity.Person
	id, _ := strconv.ParseInt(key, 10, 64)
	databases.Connector.Where("id = ?", id).Delete(&person)
	w.WriteHeader(http.StatusNoContent)
}

func UpdatePersonByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var person entity.Person
	json.Unmarshal(requestBody, &person)
	databases.Connector.Save(&person)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(person)
}

//DeletPersonByID delete's person with specific ID
func DeletePersonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var person entity.Person
	id, _ := strconv.ParseInt(key, 10, 64)
	databases.Connector.Where("id = ?", id).Delete(&person)
	w.WriteHeader(http.StatusNoContent)
}
