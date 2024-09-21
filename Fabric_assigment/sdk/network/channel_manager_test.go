package main

import (
	"testing"

	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fabric"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func TestCreateChannel(t *testing.T) {
	cm := &ChannelManager{}
	ctx := fabsdk.NewTransactionContext("mychannel", "mycc")
	err := cm.CreateChannel(ctx, "mychannel")
	if err != nil {
		t.Errorf("Error creating channel: %s", err)
		return
	}
}

func TestJoinChannel(t *testing.T) {
	cm := &ChannelManager{}
	ctx := fabsdk.NewTransactionContext("mychannel", "mycc")
	err := cm.JoinChannel(ctx, "mychannel")
	if err != nil {
		t.Errorf("Error joining channel: %s", err)
		return
	}
}

func TestGetChannel(t *testing.T) {
	cm := &ChannelManager{}
	ctx := fabsdk.NewTransactionContext("mychannel", "mycc")
	_, err := cm.GetChannel(ctx, "mychannel")
	if err != nil {
		t.Errorf("Error getting channel: %s", err)
		return
	}
}