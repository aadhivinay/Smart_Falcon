package main

import (
	"testing"

	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fabric"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func TestCreateAccount(t *testing.T) {
	fg := &FabricGateway{}
	err := fg.CreateAccount("dealer1", "msisdn1", "mpin1", 0, "active", 0, "", "")
	if err != nil {
		t.Errorf("Error creating account through gateway: %s", err)
		return
	}
}

func TestUpdateAccount(t *testing.T) {
	fg := &FabricGateway{}
	err := fg.UpdateAccount("dealer1", 100, 50, "credit", "remarks1")
	if err != nil {
		t.Errorf("Error updating account through gateway: %s", err)
		return
	}
}

func TestGetAccount(t *testing.T) {
	fg := &FabricGateway{}
	_, err := fg.GetAccount("dealer1")
	if err != nil {
		t.Errorf("Error getting account through gateway: %s", err)
		return
	}
}

func TestGetHistory(t *testing.T) {
	fg := &FabricGateway{}
	_, err := fg.GetHistory("dealer1")
	if err != nil {
		t.Errorf("Error getting history through gateway: %s", err)
		return
	}
}