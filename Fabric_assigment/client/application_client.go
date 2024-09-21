package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ApplicationClient struct{}

func (ac *ApplicationClient) CreateAccount(dealerID string, msisdn string, mpin string, balance int, status string, transAmount int, transType string, remarks string) error {
	fmt.Println("Creating account through application client...")
	// Implement account creation logic here
	return nil
}

func (ac *ApplicationClient) UpdateAccount(dealerID string, balance int, transAmount int, transType string, remarks string) error {
	fmt.Println("Updating account through application client...")
	// Implement account update logic here
	return nil
}

func (ac *ApplicationClient) GetAccount(dealerID string) ([]byte, error) {
	fmt.Println("Getting account through application client...")
	// Implement account retrieval logic here
	return nil, nil
}

func (ac *ApplicationClient) GetHistory(dealerID string) ([]byte, error) {
	fmt.Println("Getting history through application client...")
	// Implement history retrieval logic here
	return nil, nil
}

func main() {
	ac := &ApplicationClient{}
	r := mux.NewRouter()
	r.HandleFunc("/account", ac.CreateAccount).Methods("POST")
	r.HandleFunc("/account/{dealerID}", ac.UpdateAccount).Methods("PUT")
	r.HandleFunc("/account/{dealerID}", ac.GetAccount).Methods("GET")
	r.HandleFunc("/history/{dealerID}", ac.GetHistory).Methods("GET")
	http.ListenAndServe(":8080", r)
}