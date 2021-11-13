package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)



func initializeRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/",Hello).Methods("GET")
	r.HandleFunc("/listing_alert", FetchNewListing).Methods( "POST")

	log.Fatal(http.ListenAndServe("0.0.0.0:1000", r))
	log.Println("Running server on port 1000")
}

func main(){
	initializeRouter()
}

