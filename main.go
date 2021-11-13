package main

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
)


var DB *gorm.DB
var err error

func initializeRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/",Hello).Methods("GET")
	r.HandleFunc("/listing_alert", CreateListing).Methods( "POST")
	r.HandleFunc("/listing_alert", GetListings).Methods( "GET")

	log.Fatal(http.ListenAndServe("0.0.0.0:1000", r))
	log.Println("Running server on port 1000")
}

func main(){
	InitialMigration()
	initializeRouter()
}

