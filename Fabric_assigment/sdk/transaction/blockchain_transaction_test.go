package main

import (
	"testing"

	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fabric"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func TestCreateTransaction(t *testing.T) {
	bt := &BlockchainTransaction{}
	ctx := fabsdk.NewTransactionContext("mychannel", "mycc")
	err := bt.CreateTransaction(ctx, "dealer1", "msisdn1", "mpin1", 0, "active", 0, "", "")
	if err != nil {
		t.Errorf("Error creating transaction: %s", err)
		return
	}
}

func TestGetTransaction(t *testing.T) {
	bt := &BlockchainTransaction{}
	ctx := fabsdk.NewTransactionContext("mychannel", "mycc")
	_, err := bt.GetTransaction(ctx, "dealer1")
	if err != nil {
		t.Errorf("Error getting transaction: %s", err)
		return
	}
}