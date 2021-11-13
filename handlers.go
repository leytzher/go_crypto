package main

import (
	"encoding/json"
	"net/http"
)


func GetListings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var listings []CoinListing
	DB.Find(&listings)
	json.NewEncoder(w).Encode(listings)
}

func Hello(w http.ResponseWriter, r *http.Request){
	listing := CoinListing{
		Type:     "test",
		Message:  "This is a test",
		Currency: "NONE",
		Exchange: "NONE",
	}
	json.NewEncoder(w).Encode(listing)
}

func CreateListing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var listing CoinListing
	json.NewDecoder(r.Body).Decode(&listing)
	DB.Create(&listing)
	json.NewEncoder(w).Encode(listing)
}