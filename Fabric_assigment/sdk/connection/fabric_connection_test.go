package main

import (
	"testing"

	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fabric"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func TestConnect(t *testing.T) {
	sdk, err := fabsdk.New(config.FromFile("fabric_sdk_config.json"))
	if err != nil {
		t.Errorf("Error creating SDK: %s", err)
		return
	}

	defer sdk.Close()

	client, err := sdk.Client()
	if err != nil {
		t.Errorf("Error getting client: %s", err)
		return
	}

	if client == nil {
		t.Errorf("Client is nil")
		return
	}
}