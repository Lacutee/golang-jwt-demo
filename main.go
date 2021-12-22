package main

import (
	"golang-jwt-demo/database"
	"golang-jwt-demo/models"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var router *mux.Router

func CreateRouter() {
	router = mux.NewRouter()
}

func InitializeRouter() {
	router.HandleFunc("/SignUp", SignUp).Methods("POST")
	router.HandleFunc("/SignIn", SignIn).Methods("POST")
}

func initDB() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "",
			DB:         "go_lang_auth",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.UserMigrate(&models.User{})
	database.AuthMigrate(&models.Authentication{})
	database.TokenMigrate(&models.Token{})
}

func main() {
	initDB()
	CreateRouter()
	InitializeRouter()
}
