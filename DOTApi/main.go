package main

import (
	"DOTApi/services"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
	_, err := fmt.Fprintf(w, "Welcome to the DOTScanner Api!")
	if err != nil {
		return
	}
	fmt.Println("Welcome to the DOTScanner Api!")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)

	putRequest := myRouter.Methods(http.MethodPut).Subrouter()
	putRequest.HandleFunc("/api/v1/users", services.RefreshToken).Methods("PUT")

	postRequest := myRouter.Methods(http.MethodPost).Subrouter()
	postRequest.HandleFunc("/api/v1/users", services.CreateNewUser).Methods("POST")
	postRequest.HandleFunc("/api/v1/scans", services.CreateNewScan).Methods("POST")

	getRequest := myRouter.Methods(http.MethodGet).Subrouter()
	getRequest.HandleFunc("/api/v1/scans", services.ReturnAllScans).Methods("GET")
	getRequest.HandleFunc("/api/v1/scans/{id}", services.ReturnSingleScan).Methods("GET")
	getRequest.HandleFunc("/api/v1/scan-types", services.ReturnAllScanTypes).Methods("GET")
	getRequest.HandleFunc("/api/v1/scan-types/{id}", services.ReturnSingleScanType).Methods("GET")
	getRequest.HandleFunc("/api/v1/users/{email}/{password}", services.ReturnSingleUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequests()
}