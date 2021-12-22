package routers

import (
	"encoding/json"
	"fmt"
	"golang-jwt-demo/_helper"
	"golang-jwt-demo/database"
	"golang-jwt-demo/models"
	"io/ioutil"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)

	var user models.User
	var dbUser models.User

	json.Unmarshal(requestBody, &user)

	// error from body json
	err := json.NewDecoder(r.Body).Decode(&user)

	//  check if email exist
	database.Connector.Where("email = ?", user.Email).First(&dbUser)
	if dbUser.Email != "" {
		fmt.Println("Email has already taken!")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Email has already taken!")
	}

	// hashing password
	user.Password, err = _helper.GenerateHashPassword(user.Password)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error while hashing Password")
	}

	errDB := database.Connector.Create(&user).Error

	if errDB != nil {
		fmt.Println(errDB)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errDB)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func SigIn(w http.ResponseWriter, r *http.Request) {

}
