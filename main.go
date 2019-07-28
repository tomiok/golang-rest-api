package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"rest-store/model"
	"strconv"
)

const port = ":8080"

func main() {

	log.Print("From log")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/computers", addComputerHandler).Methods(http.MethodPost)
	router.HandleFunc("/computers", getAll).Methods(http.MethodGet)
	router.HandleFunc("/computers/{serial}", getBySerial).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(port, router))
}

func getBySerial(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	serial := mux.Vars(request)["serial"]

	i, _ := strconv.Atoi(serial)
	c, _ := model.SearchBySerial(int64(i))

	_ = json.NewEncoder(writer).Encode(c)
}

func getAll(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(writer).Encode(model.GetAllComputers())
}

func addComputerHandler(w http.ResponseWriter, r *http.Request) {
	var pc model.Computer
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &pc)
	model.Computers = append(model.Computers, pc)
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(pc)
}
