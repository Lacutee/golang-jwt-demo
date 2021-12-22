package routers

import (
	"encoding/json"
	"fmt"
	"golang-jwt-demo/database"
	"golang-jwt-demo/models"
	"net/http"
)

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	var user []models.User

	err := database.Connector.Find(&user).Error

	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}
