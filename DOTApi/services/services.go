package services

import (
	"DOTApi/crypto"
	"DOTApi/dal"
	"DOTApi/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var Users []models.User

func ReturnAllScans(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	//log.Println(r.Header.Get("Date"))

	var userIdRequest = r.URL.Query().Get("user-id")
	var userId, _ = strconv.ParseInt(userIdRequest, 10, 64)

	if userIdRequest != "" {
		err := json.NewEncoder(w).Encode(dal.GetAllScansByUserId(userId))
		if err != nil {
			return
		}
	} else {
		err := json.NewEncoder(w).Encode(dal.GetAllScans())
		if err != nil {
			return
		}
	}
}

func ReturnAllScanTypes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(dal.GetAllScanTypes())
	if err != nil {
		return
	}
}

func ReturnSingleScan(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	key, _ := strconv.ParseInt(vars["id"], 10 , 64)

	err := json.NewEncoder(w).Encode(dal.GetScanById(key))
	if err != nil {
		return
	}
}

func ReturnSingleScanType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	key, _ := strconv.ParseInt(vars["id"], 10 , 64)

	err := json.NewEncoder(w).Encode(dal.GetScanTypeById(key))
	if err != nil {
		return
	}
}

func ReturnSingleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	email, _ := vars["email"]
	//password, _ := vars["password"]

	for _, user := range Users {
		if user.Email == email {
			err := json.NewEncoder(w).Encode(Users)
			if err != nil {
				return
			}
		}
	}
}

func CreateNewScan(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var scan models.Scan

	err := json.Unmarshal(reqBody, &scan)
	if err != nil {
		return
	}

	dal.InsertScan(scan)
}

func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var user models.User

	err := json.Unmarshal(reqBody, &user)
	if err != nil {
		panic(err.Error())
	}

	// Encrypt the user password
	user.Password = crypto.Encrypt(user.Password)

	dal.InsertUser(user)

	log.Println(crypto.Decrypt(user.Password))
}