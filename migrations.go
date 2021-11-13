package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CoinListing struct {
	gorm.Model
	Type     string `json:"type"`
	Message  string `json:"message"`
	Currency string `json:"currency"`
	Exchange string `json:"exchange"`
}

func InitialMigration() {

	errs := godotenv.Load()
	if errs != nil {
		log.Fatal("Error loading .env file")
	}
	// Get environment variables

	user := os.Getenv("POSTGRES_USER")
	passwd := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", host, user, passwd, dbName)

	// Connect to the database
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to Database")
	}
	DB.AutoMigrate(&CoinListing{})

}
