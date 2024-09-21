package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type RESTAPI struct{}

func (ra *RESTAPI) CreateAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating account through REST API...")
	// Implement account creation logic here
}

func (ra *RESTAPI) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating account through REST API...")
	// Implement account update logic here
}

func (ra *RESTAPI) GetAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting account through REST API...")
	// Implement account retrieval logic here
}

func (ra *RESTAPI) GetHistory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting history through REST API...")
	// Implement history retrieval logic here
}

func main() {
	ra := &RESTAPI{}
	r