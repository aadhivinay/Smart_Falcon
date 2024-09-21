package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestCreateAccount(t *testing.T) {
	ac := &ApplicationClient{}
	r := mux.NewRouter()
	r.HandleFunc("/account", ac.CreateAccount).Methods("POST")
	req, err := http.NewRequest("POST", "/account", bytes.NewBuffer([]byte(`{"dealerID":"dealer1","msisdn":"msisdn1","mpin":"mpin1","balance":0,"status":"active","transAmount":0,"transType":"","remarks":""}`)))
	if err != nil {
		t.Errorf("Error creating request: %s", err)
		return
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		return
	}
}

func TestUpdateAccount(t *testing.T) {
	ac := &ApplicationClient{}
	r := mux.NewRouter()
	r.HandleFunc("/account/{dealerID}", ac.UpdateAccount).Methods("PUT")
	req, err := http.NewRequest("PUT", "/account/dealer1", bytes.NewBuffer([]byte(`{"balance":100,"transAmount":50,"transType":"credit","remarks":"remarks1"}`)))
	if err != nil {
		t.Errorf("Error