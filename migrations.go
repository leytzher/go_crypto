package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB
var err error


type CoinListing struct {
	gorm.Model
	Type     string `json:"type"`
	Message  string `json:"message"`
	Currency string `json:"currency"`
	Exchange string `json:"exchange"`
}

func InitialMigration(){
	err:= godotenv.Load()
	b := err != nil
	if b {
		log.Fatal("Error loading .env file")
	}
	user := os.Getenv("POSTGRES_USER")
	passwd := os.Getenv("POSTGRES_PASSWORD")
	db := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	dns := fmt.Sprintf("host= user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", host,user, passwd, db)
	DB, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
	fmt.Println(err.Error())
	panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&CoinListing{})

}