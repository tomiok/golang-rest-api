package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"rest-store/model"
)

func main() {

	log.Print("From log")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/computers", addComputerHandler).Methods("POST")
	router.HandleFunc("/computers", getAll).Methods("GET")
	//router.HandleFunc("/computers/{serial}", getBySerial).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getAll(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)

	_ = json.NewEncoder(writer).Encode(model.GetAllComputers())
}

func addComputerHandler(w http.ResponseWriter, r *http.Request) {
	var pc model.Computer
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &pc)
	model.Computers = append(model.Computers, pc)
	w.WriteHeader(201)
	_ = json.NewEncoder(w).Encode(pc)
}
