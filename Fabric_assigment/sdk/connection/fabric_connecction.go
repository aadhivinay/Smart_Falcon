package main

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fabric"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func main() {
	sdk, err := fabsdk.New(config.FromFile("fabric_sdk_config.json"))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer sdk.Close()

	client, err := sdk.Client()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Connected to Fabric network")
}