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

	err := database.Connector.Create(&user).Error

	//  check if email exist
	database.Connector.Where("email = ?", user.Email).First(&dbUser)

	if dbUser.Email != "" {
		fmt.Println("Email has already taken!")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Email has already taken!")
	} else {
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
		}
	}

	user.Password, err = _helper.GenerateHashPassword(user.Password)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error while hashing Password")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
