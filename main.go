package main

import (
	"fmt"
	"golang-jwt-demo/database"
	"golang-jwt-demo/middleware"
	"golang-jwt-demo/models"
	"golang-jwt-demo/routers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitializeRouter(router *mux.Router) {
	router.HandleFunc("/register", routers.SignUp).Methods("POST")
	router.HandleFunc("/login", routers.SigIn).Methods("POST")
	router.HandleFunc("/user/", middleware.Authorization(routers.GetAllUser)).Methods("GET")
	router.Methods("OPTION").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
	})
}

func CreateRouter() {
	router := mux.NewRouter().StrictSlash(true)
	fmt.Println("server running in port 8090")
	InitializeRouter(router)
	log.Fatal(http.ListenAndServe(":8090", router))
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

}
