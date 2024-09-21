package main

import (
	"testing"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
)

func TestAccountManager(t *testing.T) {
	am := &AccountManager{}
	stub := shim.NewMockStub("accountManager", am)

	// Test createAccount
	args := []string{"dealer1", "1234567890", "1234", "100", "active"}
	res := am.createAccount(stub, args)
	if res.Status != shim.OK {
		t.Errorf("createAccount failed: %s", res.Message)
	}

	// Test getAccount
	args = []string{"1234567890"}
	res = am.getAccount(stub, args)
	if res.Status != shim.OK {
		t.Errorf("getAccount failed: %s", res.Message)
	}

	// Test updateAccount
	args = []string{"dealer1", "1234567890", "1234", "200", "active"}
	res = am.updateAccount(stub, args)
	if res.Status != shim.OK {
		t.Errorf("updateAccount failed: %s", res.Message)
	}

	// Test deleteAccount
	args = []string{"1234567890"}
	res = am.deleteAccount(stub, args)
	if res.Status != shim.OK {
		t.Errorf("deleteAccount failed: %s", res.Message)
	}
}