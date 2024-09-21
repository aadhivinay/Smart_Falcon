package main

import (
	"testing"

	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fabric"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func TestCreateAccount(t *testing.T) {
	bl := &BusinessLogic{}
	ctx := fabsdk.NewTransactionContext("mychannel", "mycc")
	err := bl.CreateAccount(ctx, "dealer1", "msisdn1", "mpin1", 0, "active", 0, "", "")
	if err != nil {
		t.Errorf("Error creating account: %s", err)
		return
	}
}

func TestUpdateAccount(t *testing.T) {
	bl := &BusinessLogic{}
	ctx := fabsdk.NewTransactionContext("mychannel", "mycc")
	err := bl.UpdateAccount(ctx, "dealer1", 100, 50, "credit", "remarks1")
	if err != nil {
		t.Errorf("Error updating account: %s", err)
		return
	}
}

func TestGetAccount(t *testing.T) {
	bl := &BusinessLogic{}
	ctx := fabsdk.NewTransactionContext("mychannel", "mycc")
	_, err := bl.GetAccount(ctx, "dealer1")
	if err != nil {
		t.Errorf("Error getting account: %s", err)
		return
	}
}

func TestGetHistory(t *testing.T) {
	bl := &BusinessLogic{}
	ctx := fabsdk.NewTransactionContext("mychannel", "mycc")
	_, err := bl.GetHistory(ctx, "dealer1")
	if err != nil {
		t.Errorf("Error getting history: %s", err)
		return
	}
}