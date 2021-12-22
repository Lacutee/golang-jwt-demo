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
		return
	}

	// hashing password
	user.Password, err = _helper.GenerateHashPassword(user.Password)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error while hashing Password")
		return
	}

	errDB := database.Connector.Create(&user).Error

	if errDB != nil {
		fmt.Println(errDB)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errDB)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func SigIn(w http.ResponseWriter, r *http.Request) {
	// requestBody, _ := ioutil.ReadAll(r.Body)

	var auth models.Authentication
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&auth)
	if err != nil {
		fmt.Println("Something went wrong!", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid input!")
	}
	// fmt.Println(auth)
	database.Connector.Where("email = ?", auth.Email).First(&user)
	// fmt.Println(user.Email)
	if user.Email == "" {
		fmt.Println("Email not found!")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Username or Password is incorrect!")
		return
	}

	check := _helper.CheckPasswordHash(auth.Password, user.Password)

	if !check {
		fmt.Println("Username or Password is incorrect")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Username or Password is incorrect!")
		return
	}

	validToken, err := _helper.GenerateJWT(auth.Email, user.Role)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Failed to generate token")
		return
	}

	var token models.Token

	token.Email = user.Email
	token.Role = user.Role
	token.TokenString = validToken

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}
