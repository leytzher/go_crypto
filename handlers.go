package main

import (
	"encoding/json"
	"net/http"
)

type CoinListing struct {
	Type     string `json:"type"`
	Message  string `json:"message"`
	Currency string `json:"currency"`
	Exchange string `json:"exchange"`
}

func FetchNewListing (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var listing CoinListing
	json.NewDecoder(r.Body).Decode(&listing)
	json.NewEncoder(w).Encode(listing)
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