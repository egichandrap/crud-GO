package main

import (
	"log"
	"net/http"
	"rest-go-demo/controller"
	"rest-go-demo/databases"
	"rest-go-demo/entity"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	initDB()
	log.Println("Statrting the HTTP server on port 8090")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8090", router))
}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/create", controller.CreatePerson).Methods("POST")
	router.HandleFunc("/get", controller.GetAllPerson).Methods("GET")
	router.HandleFunc("/get/{id}", controller.GetPersonByID).Methods("GET")
	router.HandleFunc("/update/{id}", controller.UpdatePersonByID).Methods("PUT")
	router.HandleFunc("/delete/{id}", controller.DeletePersonByID).Methods("DELETE")
}

func initDB() {
	config := databases.Config{
		ServerName: "localhost:3306",
		User:       "root",
		Password:   "12345678",
		DB:         "learning",
	}

	var connectionString = databases.GetConnectionString(config)
	err := databases.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	databases.Migrate(&entity.Person{})
}
