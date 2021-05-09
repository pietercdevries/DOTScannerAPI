package services

import (
	"DOTApi/authenticate"
	"DOTApi/crypto"
	"DOTApi/dal"
	"DOTApi/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func validateToken(token string) bool {
	var userId int64

	userId = dal.GetUserIdByToken(token)

	if userId > 0 {
		return true
	} else {
		return false
	}
}

func validateRefreshToken(refreshToken string) bool {
	var userId int64

	userId = dal.GetUserIdByRefreshToken(refreshToken)

	if userId > 0 {
		return true
	} else {
		return false
	}
}

func RefreshToken(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var token = r.Header.Get("token")
	var refreshToken = r.Header.Get("refreshToken")

	if validateToken(token) == false{
		return
	}

	if validateRefreshToken(refreshToken) == false{
		return
	}

	var user models.User
	user.Id = dal.GetUserIdByToken(token)
	user.Token = authenticate.GenerateAuthenticateToken()
	user.RefreshToken = authenticate.GenerateRefreshToken()
	user.TokenExpireDate = time.Now().Add(time.Minute * 30)

	dal.UpdateUserTokens(user)

	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		return
	}
}

func ReturnAllScans(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	if validateToken(r.Header.Get("token")) == false{
		return
	}

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

	if validateToken(r.Header.Get("token")) == false{
		return
	}

	err := json.NewEncoder(w).Encode(dal.GetAllScanTypes())
	if err != nil {
		return
	}
}

func ReturnSingleScan(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	if validateToken(r.Header.Get("token")) == false{
		return
	}

	vars := mux.Vars(r)
	key, _ := strconv.ParseInt(vars["id"], 10 , 64)

	err := json.NewEncoder(w).Encode(dal.GetScanById(key))
	if err != nil {
		return
	}
}

func ReturnSingleScanType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if validateToken(r.Header.Get("token")) == false{
		return
	}

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
	password, _ := vars["password"]

	err := json.NewEncoder(w).Encode(dal.GetUserByUserNamePassword(email, password))
	if err != nil {
		return
	}
}

func CreateNewScan(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	if validateToken(r.Header.Get("token")) == false{
		return
	}

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
}