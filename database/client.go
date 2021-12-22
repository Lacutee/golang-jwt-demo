package database

import (
	"golang-jwt-demo/models"
	"log"

	"github.com/jinzhu/gorm"
)

var Connector *gorm.DB

func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Println("Connection was successfully build!")
	return nil
}

func UserMigrate(table *models.User) {
	Connector.AutoMigrate(&table)
	log.Println(" User table has been migrated")
}

func AuthMigrate(table *models.Authentication) {
	Connector.AutoMigrate(&table)
	log.Println(" Authentication table has been migrated")
}

func TokenMigrate(table *models.Token) {
	Connector.AutoMigrate(&table)
	log.Println(" Token table has been migrated")
}
