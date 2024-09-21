package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
)

type AccountManager struct {
}

func (am *AccountManager) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (am *AccountManager) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "createAccount" {
		return am.createAccount(stub, args)
	} else if function == "updateAccount" {
		return am.updateAccount(stub, args)
	} else if function == "getAccount" {
		return am.getAccount(stub, args)
	} else if function == "deleteAccount" {
		return am.deleteAccount(stub, args)
	}
	return shim.Error("Invalid function name")
}

func (am *AccountManager) createAccount(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 5 {
		return shim.Error("Invalid number of arguments")
	}
	account := Account{
		DEALERID: args[0],
		MSISDN:  args[1],
		MPIN:    args[2],
		BALANCE: args[3],
		STATUS:  args[4],
	}
	accountBytes, err := json.Marshal(account)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(account.MSISDN, accountBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (am *AccountManager) updateAccount(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 5 {
		return shim.Error("Invalid number of arguments")
	}
	account := Account{
		DEALERID: args[0],
		MSISDN:  args[1],
		MPIN:    args[2],
		BALANCE: args[3],
		STATUS:  args[4],
	}
	accountBytes, err := json.Marshal(account)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(account.MSISDN, accountBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (am *AccountManager) getAccount(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Invalid number of arguments")
	}
	accountBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	account := Account{}
	err = json.Unmarshal(accountBytes, &account)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(accountBytes)
}

func (am *AccountManager) deleteAccount(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Invalid number of arguments")
	}
	err := stub.DelState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

type Account struct {
	DEALERID string `json:"dealerId"`
	MSISDN   string `json:"msisdn"`
	MPIN     string `json:"mpin"`
	BALANCE  string `json:"balance"`
	STATUS   string `json:"status"`
}