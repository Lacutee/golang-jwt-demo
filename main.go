package main

import (
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

func main() {
	CreateRouter()
	InitializeRouter()
}
